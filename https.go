package flag

import (
	"github.com/urfave/cli/v2"
)

const (
	httpsSSLCertLocation = "https-ssl-cert-location"
	httpsSSLKeyLocation  = "https-ssl-key-location"
	httpsPort            = "https-port"
)

type HttpsOpts struct {
	HttpsSSLCertLocation string
	HttpsSSLKeyLocation  string
	HttpsPort            int
}

func NewDefaultHttpsOpts() *HttpsOpts {
	return &HttpsOpts{
		HttpsSSLCertLocation: "",
		HttpsSSLKeyLocation:  "",
		HttpsPort:            443,
	}
}

func GetHttpsOpts(c *cli.Context) *HttpsOpts {
	return &HttpsOpts{
		HttpsSSLCertLocation: c.String(httpsSSLCertLocation),
		HttpsSSLKeyLocation:  c.String(httpsSSLKeyLocation),
		HttpsPort:            c.Int(httpsPort),
	}
}

func (h *HttpsOpts) IsEnabled() bool {
	return h.HttpsSSLCertLocation != "" && h.HttpsSSLKeyLocation != ""
}

func AddHttpsFlags(flags *[]cli.Flag) {
	defaultOpts := NewDefaultHttpsOpts()
	*flags = append(*flags, &cli.StringFlag{
		Name:    httpsSSLCertLocation,
		Value:   defaultOpts.HttpsSSLCertLocation,
		Usage:   "HTTPS SSL cert `location`",
		EnvVars: []string{"HTTPS_SSL_CERT_LOCATION"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    httpsSSLKeyLocation,
		Value:   defaultOpts.HttpsSSLKeyLocation,
		Usage:   "HTTPS SSL key `location`",
		EnvVars: []string{"HTTPS_SSL_KEY_LOCATION"},
	})
	*flags = append(*flags, &cli.IntFlag{
		Name:    httpsPort,
		Value:   defaultOpts.HttpsPort,
		Usage:   "HTTPS `port`",
		EnvVars: []string{"HTTPS_PORT"},
	})
}
