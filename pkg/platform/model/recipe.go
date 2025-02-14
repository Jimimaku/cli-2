package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/multilog"
	"github.com/ActiveState/cli/internal/retryhttp"
	"github.com/ActiveState/cli/internal/rtutils/p"
	"github.com/ActiveState/cli/pkg/platform/api/inventory"
	iop "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_client/inventory_operations"
	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
	"github.com/ActiveState/cli/pkg/platform/authentication"
	"github.com/ActiveState/cli/pkg/sysinfo"
	"github.com/go-openapi/strfmt"
)

type SolverError struct {
	wrapped          error
	validationErrors []string
	isTransient      bool
}

func (e *SolverError) Error() string {
	return "solver_error"
}

func (e *SolverError) Unwrap() error {
	return e.wrapped
}

func (e *SolverError) ValidationErrors() []string {
	return e.validationErrors
}

func (e *SolverError) IsTransient() bool {
	return e.isTransient
}

// HostPlatform stores a reference to current platform
var HostPlatform string

// Recipe aliases recipe model
type Recipe = inventory_models.Recipe

// OrderAnnotations are sent with every order for analytical purposes described here:
// https://docs.google.com/document/d/1nXeNCRWX-4ULtk20t3C7kTZDCSBJchJJHhzz-lQuVsU/edit#heading=h.o93wm4bt5ul9
type OrderAnnotations struct {
	CommitID     string `json:"commit_id"`
	Project      string `json:"project"`
	Organization string `json:"organization"`
}

func init() {
	HostPlatform = sysinfo.OS().String()
	if osName, ok := os.LookupEnv(constants.OverrideOSNameEnvVarName); ok {
		HostPlatform = osName
	}
}

// FetchRawRecipeForCommit returns a recipe from a project based off a commitID
func FetchRawRecipeForCommit(commitID strfmt.UUID, owner, project string) (string, error) {
	return fetchRawRecipe(commitID, owner, project, nil)
}

// FetchRawRecipeForCommitAndPlatform returns a recipe from a project based off a commitID and platform
func FetchRawRecipeForCommitAndPlatform(commitID strfmt.UUID, owner, project string, platform string) (string, error) {
	return fetchRawRecipe(commitID, owner, project, &platform)
}

func ResolveRecipe(commitID strfmt.UUID, owner, projectName string) (*inventory_models.Recipe, error) {
	if commitID == "" {
		return nil, locale.NewError("err_no_commit", "Missing Commit ID")
	}

	recipe, err := FetchRecipe(commitID, owner, projectName, &HostPlatform)
	if err != nil {
		return nil, err
	}

	if recipe.RecipeID == nil {
		return nil, errs.New("Resulting recipe does not have a recipeID")
	}

	return recipe, nil
}

func fetchRawRecipe(commitID strfmt.UUID, owner, project string, hostPlatform *string) (string, error) {
	_, transport := inventory.Init(authentication.LegacyGet())

	var err error
	params := iop.NewResolveRecipesParams()
	params.SetHTTPClient(retryhttp.DefaultClient.StandardClient())
	params.SetTimeout(time.Second * 60)
	params.Order, err = commitToOrder(commitID, owner, project)
	if err != nil {
		return "", errs.Wrap(err, "commitToOrder failed")
	}

	if hostPlatform != nil {
		params.Order.Platforms, err = filterPlatformIDs(*hostPlatform, runtime.GOARCH, params.Order.Platforms)
		if err != nil {
			return "", err
		}
	}

	recipe, err := inventory.ResolveRecipes(transport, params, authentication.ClientAuth())
	if err != nil {
		if err == context.DeadlineExceeded {
			return "", locale.WrapError(err, "request_timed_out")
		}

		orderBody, err2 := json.Marshal(params.Order)
		if err2 != nil {
			orderBody = []byte(fmt.Sprintf("Could not marshal order, error: %v", err2))
		}

		serr := resolveSolverError(err)
		multilog.Error("Solver returned error: %s, order: %s", errs.JoinMessage(err), string(orderBody))
		return "", serr
	}

	return recipe, nil
}

func commitToOrder(commitID strfmt.UUID, owner, project string) (*inventory_models.Order, error) {
	monoOrder, err := FetchOrderFromCommit(commitID)
	if err != nil {
		return nil, locale.WrapError(err, "err_order_recipe")
	}

	orderData, err := monoOrder.MarshalBinary()
	if err != nil {
		return nil, locale.WrapError(err, "err_order_marshal")
	}

	order := &inventory_models.Order{}
	err = order.UnmarshalBinary(orderData)
	if err != nil {
		return nil, locale.WrapError(err, "err_order_marshal")
	}

	order.Annotations = OrderAnnotations{
		CommitID:     commitID.String(),
		Project:      project,
		Organization: owner,
	}

	return order, nil
}

func FetchRecipe(commitID strfmt.UUID, owner, project string, hostPlatform *string) (*inventory_models.Recipe, error) {
	var err error
	params := iop.NewResolveRecipesParams()
	params.SetHTTPClient(retryhttp.DefaultClient.StandardClient())
	params.SetTimeout(time.Second * 60)
	params.Order, err = commitToOrder(commitID, owner, project)
	if err != nil {
		return nil, errs.Wrap(err, "commitToOrder failed")
	}

	client, _ := inventory.Init(authentication.LegacyGet())

	response, err := client.ResolveRecipes(params, authentication.ClientAuth())
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, locale.WrapError(err, "request_timed_out")
		}

		orderBody, _ := json.Marshal(params.Order)
		serr := resolveSolverError(err)
		multilog.Error("Solver returned error: %s, order: %s", errs.JoinMessage(err), string(orderBody))
		return nil, serr
	}

	platformIDs, err := filterPlatformIDs(*hostPlatform, runtime.GOARCH, params.Order.Platforms)
	if err != nil {
		return nil, errs.Wrap(err, "filterPlatformIDs failed")
	}
	if len(platformIDs) == 0 {
		return nil, locale.NewInputError("err_recipe_no_platform")
	} else if len(platformIDs) > 1 {
		logging.Debug("Received multiple platform IDs.  Picking the first one.")
	}
	platformID := platformIDs[0]

	for _, recipe := range response.Payload.Recipes {
		if recipe.Platform != nil && recipe.Platform.PlatformID != nil && *recipe.Platform.PlatformID == platformID {
			return recipe, nil
		}
	}

	return nil, locale.NewInputError("err_recipe_not_found")
}

func resolveSolverError(err error) error {
	switch serr := err.(type) {
	case *iop.ResolveRecipesDefault:
		return &SolverError{
			wrapped:     locale.WrapError(errs.Wrap(err, "ResolveRecipesDefault"), "", p.PStr(serr.Payload.Message)),
			isTransient: p.PBool(serr.GetPayload().IsTransient),
		}
	case *iop.ResolveRecipesBadRequest:
		var validationErrors []string
		for _, verr := range serr.Payload.ValidationErrors {
			if verr.Error == nil {
				continue
			}
			lines := strings.Split(*verr.Error, "\n")
			validationErrors = append(validationErrors, lines...)
		}
		return &SolverError{
			wrapped:          locale.WrapError(errs.Wrap(err, "ResolveRecipesBadRequest"), "", p.PStr(serr.Payload.SolverError.Message)),
			validationErrors: validationErrors,
			isTransient:      p.PBool(serr.GetPayload().IsTransient),
		}
	default:
		return locale.WrapError(errs.Wrap(err, "unknown error"), "err_order_unknown")
	}
}

func IngredientVersionMap(recipe *inventory_models.Recipe) map[strfmt.UUID]*inventory_models.ResolvedIngredient {
	ingredientVersionMap := map[strfmt.UUID]*inventory_models.ResolvedIngredient{}

	for _, re := range recipe.ResolvedIngredients {
		if re.Ingredient.PrimaryNamespace != nil &&
			(*re.Ingredient.PrimaryNamespace == "builder" || *re.Ingredient.PrimaryNamespace == "builder-lib" || *re.Ingredient.PrimaryNamespace == NamespaceLanguage.String()) {
			continue
		}

		if re.IngredientVersion == nil || re.IngredientVersion.IngredientVersionID == nil {
			continue
		}
		ingredientVersionMap[*re.IngredientVersion.IngredientVersionID] = re
	}
	return ingredientVersionMap
}

func ArtifactMap(recipe *inventory_models.Recipe) map[strfmt.UUID]*inventory_models.ResolvedIngredient {
	artifactMap := map[strfmt.UUID]*inventory_models.ResolvedIngredient{}

	for _, re := range recipe.ResolvedIngredients {
		if re.Ingredient.PrimaryNamespace != nil &&
			(*re.Ingredient.PrimaryNamespace == "builder" || *re.Ingredient.PrimaryNamespace == "builder-lib" || *re.Ingredient.PrimaryNamespace == NamespaceLanguage.String()) {
			continue
		}

		artifactMap[re.ArtifactID] = re
	}
	return artifactMap
}

func ArtifactDescription(artifactID strfmt.UUID, artifactMap map[strfmt.UUID]*inventory_models.ResolvedIngredient) string {
	v, ok := artifactMap[artifactID]
	if !ok || v.Ingredient.Name == nil {
		return locale.Tl("unknown_artifact_description", "Artifact {{.V0}}", artifactID.String())
	}

	version := ""
	if v.IngredientVersion != nil {
		version = "@" + *v.IngredientVersion.Version
	}

	return *v.Ingredient.Name + version
}

func ParseDepTree(ingredients []*inventory_models.ResolvedIngredient, ingredientVersionMap map[strfmt.UUID]*inventory_models.ResolvedIngredient) (directdeptree map[strfmt.UUID][]strfmt.UUID, recursive map[strfmt.UUID][]strfmt.UUID) {
	directdeptree = map[strfmt.UUID][]strfmt.UUID{}
	for _, ingredient := range ingredients {
		if ingredient.IngredientVersion == nil || ingredient.IngredientVersion.IngredientVersionID == nil {
			continue
		}

		id := ingredient.IngredientVersion.IngredientVersionID
		// skip ingredients that are not mapped to artifacts
		if _, ok := ingredientVersionMap[*id]; !ok {
			continue
		}
		// Construct directdeptree entry
		if _, ok := directdeptree[*id]; !ok {
			directdeptree[*id] = []strfmt.UUID{}
		}

		// Add direct dependencies
		for _, dep := range ingredient.Dependencies {
			if dep.IngredientVersionID == nil {
				continue
			}
			// skip ingredients that are not mapped to artifacts
			if _, ok := ingredientVersionMap[*dep.IngredientVersionID]; !ok {
				continue
			}
			directdeptree[*id] = append(directdeptree[*id], *dep.IngredientVersionID)
		}
	}

	// Now resolve ALL dependencies, not just the direct ones
	deptree := map[strfmt.UUID][]strfmt.UUID{}
	for ingredientID := range directdeptree {
		deps := []strfmt.UUID{}
		deptree[ingredientID] = recursiveDeps(deps, directdeptree, ingredientID, map[strfmt.UUID]struct{}{})
	}

	return directdeptree, deptree
}

func recursiveDeps(deps []strfmt.UUID, directdeptree map[strfmt.UUID][]strfmt.UUID, id strfmt.UUID, skip map[strfmt.UUID]struct{}) []strfmt.UUID {
	if _, ok := directdeptree[id]; !ok {
		return deps
	}

	for _, dep := range directdeptree[id] {
		if _, ok := skip[dep]; ok {
			continue
		}
		skip[dep] = struct{}{}

		deps = append(deps, dep)
		deps = recursiveDeps(deps, directdeptree, dep, skip)
	}

	return deps
}

// IsPlatformError returns true if the error is a solver error due to a missing build image for the requested platform
func IsPlatformError(err error) bool {
	// todo: replace with error codes once we have a solution -- https://www.pivotaltracker.com/story/show/178865201
	serr := &SolverError{}
	if !errors.As(err, &serr) {
		return false
	}
	return strings.Contains(strings.Join(serr.ValidationErrors(), ""), "build image for platform")
}
