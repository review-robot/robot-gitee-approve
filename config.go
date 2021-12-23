package main

import "github.com/opensourceways/community-robot-lib/config"

type configuration struct {
	ConfigItems []botConfig `json:"config_items,omitempty"`
}

func (c *configuration) configFor(org, repo string) *botConfig {
	if c == nil {
		return nil
	}

	items := c.ConfigItems
	v := make([]config.IRepoFilter, len(items))
	for i := range items {
		v[i] = &items[i]
	}

	if i := config.Find(org, repo, v); i >= 0 {
		return &items[i]
	}

	return nil
}

func (c *configuration) Validate() error {
	if c == nil {
		return nil
	}

	items := c.ConfigItems
	for i := range items {
		if err := items[i].validate(); err != nil {
			return err
		}
	}

	return nil
}

func (c *configuration) SetDefault() {
	if c == nil {
		return
	}

	Items := c.ConfigItems
	for i := range Items {
		Items[i].setDefault()
	}
}

type botConfig struct {
	config.RepoFilter
	// IssueRequired indicates if an associated issue is required for approval in
	// the specified repos.
	IssueRequired bool `json:"issue_required,omitempty"`

	// RequireSelfApproval requires PR authors to explicitly approve their PRs.
	// Otherwise the plugin assumes the author of the PR approves the changes in the PR.
	RequireSelfApproval *bool `json:"require_self_approval,omitempty"`

	// LgtmActsAsApprove indicates that the lgtm command should be used to
	// indicate approval
	LgtmActsAsApprove bool `json:"lgtm_acts_as_approve,omitempty"`

	// IgnoreReviewState causes the approve plugin to ignore the gitee review state. Otherwise:
	// * an APPROVE github review is equivalent to leaving an "/approve" message.
	// * A REQUEST_CHANGES github review is equivalent to leaving an /approve cancel" message.
	IgnoreReviewState *bool `json:"ignore_review_state,omitempty"`

	// TODO(fejta): delete in June 2019
	DeprecatedImplicitSelfApprove *bool `json:"implicit_self_approve,omitempty"`

	// ReviewActsAsApprove should be replaced with its non-deprecated inverse: ignore_review_state.
	// TODO(fejta): delete in June 2019
	DeprecatedReviewActsAsApprove *bool `json:"review_acts_as_approve,omitempty"`
}

func (c *botConfig) setDefault() {
	bDefault := true

	if c.IgnoreReviewState == nil {
		c.IgnoreReviewState = &bDefault
	}

	if c.RequireSelfApproval == nil {
		c.RequireSelfApproval = &bDefault
	}
}

func (c *botConfig) validate() error {
	return c.RepoFilter.Validate()
}
