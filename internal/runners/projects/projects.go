package projects

import (
	"sort"
	"strings"

	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/multilog"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/primer"
	"github.com/ActiveState/cli/pkg/platform/authentication"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/ActiveState/cli/pkg/projectfile"
)

// Holds a union of project and organization parameters.
type projectWithOrg struct {
	Name           string   `json:"name"`
	Organization   string   `json:"organization"`
	LocalCheckouts []string `json:"local_checkouts,omitempty"`
}

type projectWithOrgs []projectWithOrg

func (o projectWithOrgs) MarshalOutput(f output.Format) interface{} {
	type projectOutputPlain struct {
		Name           string
		Organization   string
		LocalCheckouts string `locale:"local_checkouts,Local Checkouts" opts:"emptyNil,separateLine"`
	}

	if f != output.PlainFormatName {
		return o
	}

	r := []projectOutputPlain{}
	for _, v := range o {
		checkouts := []string{}
		for _, checkout := range v.LocalCheckouts {
			checkouts = append(checkouts, locale.Tl("projects_local_checkout", " └─ ✔ Local Checkout → {{.V0}}", checkout))
		}
		r = append(r, projectOutputPlain{v.Name, v.Organization, strings.Join(checkouts, "\n")})
	}
	return r
}

type configGetter interface {
	projectfile.ConfigGetter
}

// Params are command line parameters
type Params struct {
	Local bool // Whether to show locally checked out projects only
}

type Projects struct {
	auth   *authentication.Auth
	out    output.Outputer
	config configGetter
}

type primeable interface {
	primer.Auther
	primer.Outputer
	primer.Configurer
}

func NewParams() *Params {
	return &Params{Local: false}
}

func NewProjects(prime primeable) *Projects {
	return newProjects(prime.Auth(), prime.Output(), prime.Config())
}

func newProjects(auth *authentication.Auth, out output.Outputer, config configGetter) *Projects {
	return &Projects{
		auth,
		out,
		config,
	}
}

func (r *Projects) Run(params *Params) error {
	localProjects := projectfile.GetProjectMapping(r.config)

	var projects projectWithOrgs = []projectWithOrg{}
	for namespace, checkouts := range localProjects {
		ns, err := project.ParseNamespace(namespace)
		if err != nil {
			multilog.Error("Invalid project namespace stored to config mapping: %s", namespace)
			continue
		}

		projects = append(projects, projectWithOrg{
			Name:           ns.Project,
			Organization:   ns.Owner,
			LocalCheckouts: checkouts,
		})
	}
	sort.SliceStable(projects, func(i, j int) bool {
		if projects[i].Organization == projects[j].Organization {
			return projects[i].Name < projects[j].Name
		}
		return projects[i].Organization < projects[j].Organization
	})

	if len(projects) == 0 {
		r.out.Print(locale.T("project_checkout_empty"))
	} else {
		r.out.Print(projects)
	}

	return nil
}
