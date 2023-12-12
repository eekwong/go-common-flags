package flag

import (
	"github.com/urfave/cli/v2"
)

const (
	githubHost = "github-host"
	githubPAT  = "github-pat"
)

type GitHubOpts struct {
	Host string
	PAT  string
}

// NewDefaultGitHubOpts returns a new GitHubOpts with default values
func NewDefaultGitHubOpts() *GitHubOpts {
	return &GitHubOpts{
		Host: "github.com",
		PAT:  "",
	}
}

// GetGitHubOpts parses the cobra.Command and returns the GitHubOpts.
func GetGitHubOpts(c *cli.Context) *GitHubOpts {
	return &GitHubOpts{
		Host: c.String(githubHost),
		PAT:  c.String(githubPAT),
	}
}

// AddGitHubFlags adds the GitHub-specific command line arguments to the cobra.Command.
func AddGitHubFlags(flags *[]cli.Flag) {
	defaultOpts := NewDefaultGitHubOpts()
	*flags = append(*flags, &cli.StringFlag{
		Name:    githubHost,
		Value:   defaultOpts.Host,
		Usage:   "GitHub `host`",
		EnvVars: []string{"GITHUB_HOST"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    githubPAT,
		Value:   defaultOpts.PAT,
		Usage:   "GitHub `PAT`",
		EnvVars: []string{"GITHUB_PAT", "GH_TOKEN"},
	})
}
