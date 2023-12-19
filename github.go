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

func GetGitHubOpts(c *cli.Context) *GitHubOpts {
	return &GitHubOpts{
		Host: c.String(githubHost),
		PAT:  c.String(githubPAT),
	}
}

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
