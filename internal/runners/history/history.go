package history

import (
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/primer"
	"github.com/ActiveState/cli/pkg/cmdlets/commit"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/go-openapi/strfmt"
)

type primeable interface {
	primer.Projecter
	primer.Outputer
}

type History struct {
	project *project.Project
	out     output.Outputer
}

func NewHistory(prime primeable) *History {
	return &History{
		prime.Project(),
		prime.Output(),
	}
}

type HistoryParams struct {
}

func (h *History) Run(params *HistoryParams) error {
	if h.project == nil {
		return locale.NewInputError("err_history_no_project", "No project found. Please run this command in a project directory")
	}

	localCommitID := h.project.CommitUUID()

	var latestRemoteID *strfmt.UUID
	if !h.project.IsHeadless() {
		remoteBranch, err := model.BranchForProjectNameByName(h.project.Owner(), h.project.Name(), h.project.BranchName())
		if err != nil {
			return locale.WrapError(err, "err_history_remote_branch", "Could not get branch by local branch name")
		}

		latestRemoteID, err = model.CommonParent(remoteBranch.CommitID, &localCommitID)
		if err != nil {
			return locale.WrapError(err, "err_history_common_parent", "Could not determine common parent commit")
		}
	}

	commits, err := model.CommitHistoryFromID(localCommitID)
	if err != nil {
		return locale.WrapError(err, "err_commit_history_commit_id", "Could not get commit history from commit ID.")
	}

	if len(commits) == 0 {
		h.out.Print(locale.Tr("no_commits", h.project.Namespace().String()))
		return nil
	}

	authorIDs := authorIDsForCommits(commits)
	orgs, err := model.FetchOrganizationsByIDs(authorIDs)
	if err != nil {
		return err
	}

	h.out.Notice(locale.Tl("history_recent_changes", "Here are the most recent changes made to this project.\n"))
	err = commit.PrintCommits(h.out, commits, orgs, latestRemoteID)
	if err != nil {
		return locale.WrapError(err, "err_history_print_commits", "Could not print commit history")
	}

	return nil
}

func authorIDsForCommits(commits []*mono_models.Commit) []strfmt.UUID {
	authorIDs := []strfmt.UUID{}
	for _, commit := range commits {
		if commit.Author != nil {
			authorIDs = append(authorIDs, *commit.Author)
		}
	}
	return authorIDs
}
