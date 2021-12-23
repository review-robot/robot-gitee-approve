package main

import (
	"context"

	"github.com/opensourceways/repo-owners-cache/grpc/client"
	"github.com/opensourceways/repo-owners-cache/protocol"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/test-infra/prow/github"
)

type ghclient struct {
	cli iClient
}

func (c *ghclient) GetPullRequestChanges(org, repo string, number int) ([]github.PullRequestChange, error) {
	cs, err := c.cli.GetPullRequestChanges(org, repo, int32(number))
	if err != nil {
		return nil, err
	}

	return transformPRChanges(cs), nil
}

func (c *ghclient) GetIssueLabels(org, repo string, number int) ([]github.Label, error) {
	labels, err := c.cli.GetPRLabels(org, repo, int32(number))
	if err != nil {
		return nil, err
	}

	return transformLabels(labels), nil
}

func (c *ghclient) ListIssueComments(org, repo string, number int) ([]github.IssueComment, error) {
	comments, err := c.cli.ListPRComments(org, repo, int32(number))
	if err != nil {
		return nil, err
	}

	return transformComments(comments), nil
}

func (c *ghclient) DeleteComment(org, repo string, ID int) error {
	return c.cli.DeletePRComment(org, repo, int32(ID))
}

func (c *ghclient) CreateComment(org, repo string, number int, comment string) error {
	return c.cli.CreatePRComment(org, repo, int32(number), comment)
}

func (c *ghclient) BotName() (string, error) {
	bot, err := c.cli.GetBot()
	if err != nil {
		return "", err
	}
	return bot.Login, nil
}

func (c *ghclient) AddLabel(org, repo string, number int, label string) error {
	return c.cli.AddPRLabel(org, repo, int32(number), label)
}

func (c *ghclient) RemoveLabel(org, repo string, number int, label string) error {
	return c.cli.RemovePRLabel(org, repo, int32(number), label)
}

func (c *ghclient) ListIssueEvents(org, repo string, num int) ([]github.ListedIssueEvent, error) {
	return []github.ListedIssueEvent{}, nil
}

func (c *ghclient) GetPullRequest(org, repo string, number int) (*github.PullRequest, error) {
	return nil, nil
}

func (c *ghclient) ListReviews(org, repo string, number int) ([]github.Review, error) {
	return []github.Review{}, nil
}

func (c *ghclient) ListPullRequestComments(org, repo string, number int) ([]github.ReviewComment, error) {
	return []github.ReviewComment{}, nil
}

func newGHClient(cli iClient) *ghclient {
	return &ghclient{cli: cli}
}

type ownersClient struct {
	cli *client.Client
	log *logrus.Entry

	org    string
	repo   string
	branch string
}

func (oc *ownersClient) genRepoFilePathParam(path string) *protocol.RepoFilePath {
	return &protocol.RepoFilePath{
		Branch: &protocol.Branch{
			Platform: "gitee",
			Org:      oc.org,
			Repo:     oc.org,
			Branch:   oc.branch,
		},
		File: path,
	}
}

func (oc *ownersClient) Approvers(path string) sets.String {
	res := sets.NewString()

	o, err := oc.cli.Approvers(context.Background(), oc.genRepoFilePathParam(path))
	if err != nil {
		oc.log.Error(err)
		return res
	}

	return res.Insert(o.GetOwners()...)
}

func (oc *ownersClient) LeafApprovers(path string) sets.String {
	res := sets.NewString()

	o, err := oc.cli.LeafApprovers(context.Background(), oc.genRepoFilePathParam(path))
	if err != nil {
		oc.log.Error(err)

		return res
	}

	return res.Insert(o.GetOwners()...)
}
func (oc *ownersClient) FindApproverOwnersForFile(file string) string {
	p, err := oc.cli.FindApproverOwnersForFile(context.Background(), oc.genRepoFilePathParam(file))
	if err != nil {
		oc.log.Error(err)

		return ""
	}

	return p.GetPath()
}

func (oc *ownersClient) IsNoParentOwners(path string) bool {
	// TODO: need upstream export grpc api
	return false
}

func newOwnersClient(cli *client.Client, log *logrus.Entry, org, repo, branch string, ) *ownersClient {
	return &ownersClient{cli: cli, log: log, org: org, repo: repo, branch: branch}
}
