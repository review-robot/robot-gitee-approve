package config

import "net/url"

// GitHubOptions allows users to control how prow applications display GitHub website links.
type GitHubOptions struct {
	// LinkURLFromConfig is the string representation of the link_url config parameter.
	// This config parameter allows users to override the default GitHub link url for all plugins.
	// If this option is not set, we assume "https://github.com".
	LinkURLFromConfig string `json:"link_url,omitempty"`

	// LinkURL is the url representation of LinkURLFromConfig. This variable should be used
	// in all places internally.
	LinkURL *url.URL `json:"-"`
}
