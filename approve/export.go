package approve

import (
	"fmt"
	"regexp"

	"k8s.io/test-infra/prow/github"
)

func NewState(org, repo, branch, body, author, url string, number int, assignees []github.User) *state {
	return &state{
		org:       org,
		repo:      repo,
		branch:    branch,
		number:    number,
		body:      body,
		author:    author,
		assignees: assignees,
		htmlURL:   url,
	}
}

var (
	Handle = handle
)

func GetBotCommandLink(url string) string {
	platform := parsePlatform(url)

	p := ""
	switch platform {
	case "gitee":
		p = "gitee-deck/"
	}

	return fmt.Sprintf("https://prow.osinfra.cn/%scommand-help", p)
}

func parsePlatform(url string) string {
	re := regexp.MustCompile(".*/(.*).com/")
	m := re.FindStringSubmatch(url)
	if m != nil {
		return m[1]
	}
	return ""
}
