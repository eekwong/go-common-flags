package flag

import (
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

const (
	oktaClientID          = "okta-client-id"
	oktaClientIDFile      = "okta-client-id-file"
	oktaClientSecret      = "okta-client-secret"
	oktaClientSecretFile  = "okta-client-secret-file"
	oktaIssuer            = "okta-issuer"
	oktaRedirectURL       = "okta-redirect-url"
	oktaLogoutRedirectURL = "okta-logout-redirect-url"
	oktaAPIPath           = "okta-api-path"
)

// OktaOpts is the options for accessing Okta
type OktaOpts struct {
	OktaClientID          string
	OktaClientIDFile      string
	OktaClientSecret      string
	OktaClientSecretFile  string
	OktaIssuer            string
	OktaRedirectURL       string
	OktaLogoutRedirectURL string
	OktaAPIPath           string
}

// NewDefaultOktaOpts returns a new OktaOpts with default values
func NewDefaultOktaOpts() *OktaOpts {
	return &OktaOpts{
		OktaClientID:          "",
		OktaClientIDFile:      "",
		OktaClientSecret:      "",
		OktaClientSecretFile:  "",
		OktaIssuer:            "",
		OktaRedirectURL:       "",
		OktaLogoutRedirectURL: "",
		OktaAPIPath:           "/oauth2",
	}
}

// GetOktaOpts parses the cobra.Command and returns the OktaOpts.
func GetOktaOpts(c *cli.Context) *OktaOpts {
	o := &OktaOpts{
		OktaClientID:          c.String(oktaClientID),
		OktaClientIDFile:      c.String(oktaClientIDFile),
		OktaClientSecret:      c.String(oktaClientSecret),
		OktaClientSecretFile:  c.String(oktaClientSecretFile),
		OktaIssuer:            c.String(oktaIssuer),
		OktaRedirectURL:       c.String(oktaRedirectURL),
		OktaLogoutRedirectURL: c.String(oktaLogoutRedirectURL),
		OktaAPIPath:           c.String(oktaAPIPath),
	}
	if o.OktaClientID == "" && o.OktaClientIDFile != "" {
		if b, err := os.ReadFile(o.OktaClientIDFile); err == nil {
			o.OktaClientID = strings.TrimSuffix(string(b), "\n")
		}
	}
	if o.OktaClientSecret == "" && o.OktaClientSecretFile != "" {
		if b, err := os.ReadFile(o.OktaClientSecretFile); err == nil {
			o.OktaClientSecret = strings.TrimSuffix(string(b), "\n")
		}
	}
	return o
}

// IsEnabled is true if all the arguments are non-empty
func (o *OktaOpts) IsEnabled() bool {
	return o.OktaClientID != "" &&
		o.OktaClientSecret != "" &&
		o.OktaIssuer != "" &&
		o.OktaRedirectURL != "" &&
		o.OktaLogoutRedirectURL != ""
}

// AddOktaFlags adds the Okta-specific command line arguments to the cobra.Command.
func AddOktaFlags(flags *[]cli.Flag) {
	defaultOpts := NewDefaultOktaOpts()
	*flags = append(*flags, &cli.StringFlag{
		Name:    oktaClientID,
		Value:   defaultOpts.OktaClientID,
		Usage:   "Okta Client ID (override file)",
		EnvVars: []string{"OKTA_CLIENT_ID"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    oktaClientIDFile,
		Value:   defaultOpts.OktaClientIDFile,
		Usage:   "Okta Client ID File",
		EnvVars: []string{"OKTA_CLIENT_ID_FILE"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    oktaClientSecret,
		Value:   defaultOpts.OktaClientSecret,
		Usage:   "Okta Client Secret (override file)",
		EnvVars: []string{"OKTA_CLIENT_SECRET"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    oktaClientSecretFile,
		Value:   defaultOpts.OktaClientSecretFile,
		Usage:   "Okta Client Secret File",
		EnvVars: []string{"OKTA_CLIENT_SECRET_FILE"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    oktaIssuer,
		Value:   defaultOpts.OktaIssuer,
		Usage:   "Okta Issuer",
		EnvVars: []string{"OKTA_ISSUER"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    oktaRedirectURL,
		Value:   defaultOpts.OktaRedirectURL,
		Usage:   "Okta Redirect URL",
		EnvVars: []string{"OKTA_REDIRECT_URL"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    oktaLogoutRedirectURL,
		Value:   defaultOpts.OktaLogoutRedirectURL,
		Usage:   "Okta Logout Redirect URL",
		EnvVars: []string{"OKTA_LOGOUT_REDIRECT_URL"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    oktaAPIPath,
		Value:   defaultOpts.OktaAPIPath,
		Usage:   "Okta API path (before /v1/*)",
		EnvVars: []string{"OKTA_API_PATH"},
	})
}
