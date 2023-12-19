package flag

import (
	"github.com/urfave/cli/v2"
)

const (
	httpPort = "http-port"
)

type HttpOpts struct {
	HttpPort int
}

func NewDefaultHttpOpts() *HttpOpts {
	return &HttpOpts{
		HttpPort: 8080,
	}
}

func GetHttpOpts(c *cli.Context) *HttpOpts {
	return &HttpOpts{
		HttpPort: c.Int(httpPort),
	}
}

func AddHttpFlags(flags *[]cli.Flag) {
	defaultOpts := NewDefaultHttpOpts()
	*flags = append(*flags, &cli.IntFlag{
		Name:    httpPort,
		Value:   defaultOpts.HttpPort,
		Usage:   "HTTP `port`",
		EnvVars: []string{"HTTP_PORT"},
	})
}
