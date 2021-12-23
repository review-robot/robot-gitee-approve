package main

import (
	"net/url"
	"regexp"
	"strings"

	sdk "github.com/opensourceways/go-gitee/gitee"
	"github.com/sirupsen/logrus"
	"k8s.io/test-infra/prow/github"

	"github.com/opensourceways/robot-gitee-approve/approve"
	"github.com/opensourceways/robot-gitee-approve/approve/config"
)

const (
	approveCommand = "APPROVE"
	lgtmCommand    = "LGTM"
)

func (bot *robot) handle(org, repo string, pr *sdk.PullRequestHook, cfg *botConfig, log *logrus.Entry) error {
	c := transformConfig(cfg)
	oc := newOwnersClient(bot.cacheCli, log, org, repo, pr.GetBase().GetRef())
	ghc := newGHClient(bot.cli)
	assignees := make([]github.User, 0, len(pr.Assignees))

	for _, v := range pr.Assignees {
		assignees = append(assignees, github.User{Login: v.GetLogin()})
	}

	state := approve.NewState(
		org, repo,
		pr.GetBase().GetRef(),
		pr.GetBody(),
		pr.GetUser().GetLogin(),
		pr.GetHtmlURL(),
		int(pr.GetNumber()),
		assignees,
	)

	return approve.Handle(log, ghc, oc, getGiteeOption(), c, state)
}

func (bot *robot) authorIsRobot(author string) (bool, error) {
	b, err := bot.cli.GetBot()
	if err != nil {
		return false, err
	}

	return b.Name == author, err
}

func isApproveCommand(comment string, lgtmActsAsApprove bool) bool {
	reg := regexp.MustCompile(`(?m)^/([^\s]+)[\t ]*([^\n\r]*)`)

	for _, match := range reg.FindAllStringSubmatch(comment, -1) {
		cmd := strings.ToUpper(match[1])
		if (cmd == lgtmCommand && lgtmActsAsApprove) || cmd == approveCommand {
			return true
		}
	}

	return false
}

func getGiteeOption() config.GitHubOptions {
	s := "https://gitee.com"
	linkURL, _ := url.Parse(s)

	return config.GitHubOptions{LinkURLFromConfig: s, LinkURL: linkURL}
}
