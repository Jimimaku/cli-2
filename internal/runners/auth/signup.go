package auth

import (
	"github.com/ActiveState/cli/internal/keypairs"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/prompt"
	authlet "github.com/ActiveState/cli/pkg/cmdlets/auth"
	"github.com/ActiveState/cli/pkg/platform/authentication"
)

type Signup struct {
	output.Outputer
	prompt.Prompter
	keypairs.Configurable
	*authentication.Auth
}

func NewSignup(prime primeable) *Signup {
	return &Signup{prime.Output(), prime.Prompt(), prime.Config(), prime.Auth()}
}

func (s *Signup) Run(params *SignupParams) error {
	if s.Auth.Authenticated() {
		return locale.NewInputError("err_auth_authenticated", "You are already authenticated as: {{.V0}}. You can log out by running `state auth logout`.", s.Auth.WhoAmI())
	}

	if !params.Prompt {
		return authlet.AuthenticateWithBrowser(s.Outputer, s.Auth, s.Prompter) // user can sign up from this page too
	}
	return authlet.Signup(s.Configurable, s.Outputer, s.Prompter, s.Auth)
}
