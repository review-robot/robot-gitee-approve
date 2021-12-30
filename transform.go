package main

import (
	"time"

	sdk "github.com/opensourceways/go-gitee/gitee"
	"k8s.io/test-infra/prow/github"

	"github.com/opensourceways/robot-gitee-approve/approve/plugins"
)

func transformPRChanges(changes []sdk.PullRequestFiles) []github.PullRequestChange {
	var res []github.PullRequestChange

	for _, v := range changes {
		res = append(res, github.PullRequestChange{
			SHA:      v.Sha,
			Filename: v.Filename,
			Status:   v.Status,
		})
	}

	return res
}

func transformLabels(labels []sdk.Label) []github.Label {
	var res []github.Label

	for _, v := range labels {
		res = append(res, github.Label{
			URL:   v.Url,
			Name:  v.Name,
			Color: v.Color,
		})
	}

	return res
}

func transformComments(comments []sdk.PullRequestComments) []github.IssueComment {
	var res []github.IssueComment

	parseTime := func(t string) time.Time {
		r, _ := time.Parse(time.RFC3339, t)

		return r
	}

	for _, v := range comments {
		res = append(res, github.IssueComment{
			ID:        int(v.Id),
			Body:      v.Body,
			User:      transformUser(v.User),
			HTMLURL:   v.HtmlUrl,
			CreatedAt: parseTime(v.CreatedAt),
			UpdatedAt: parseTime(v.UpdatedAt),
		})
	}

	return res
}

func transformUser(user *sdk.UserBasic) github.User {
	return github.User{
		Login:   user.GetLogin(),
		Name:    user.GetName(),
		Email:   user.GetEmail(),
		ID:      int(user.GetID()),
		HTMLURL: user.GetHtmlUrl(),
		Type:    user.GetType(),
	}
}

func transformConfig(org string, cfg *botConfig) plugins.Approve {
	return plugins.Approve{
		Repos:               []string{org},
		RequireSelfApproval: &cfg.RequireSelfApproval,
		IgnoreReviewState:   &cfg.ignoreReviewState,
	}
}
