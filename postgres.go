package flag

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const (
	postgresHost     = "psql-host"
	postgresPort     = "psql-port"
	postgresUsername = "psql-username"
	postgresPassword = "psql-password"
	postgresDBName   = "psql-dbname"
)

// PostgresOpts is the options for accessing PostgreSQL
type PostgresOpts struct {
	PostgresHost     string
	PostgresPort     int
	PostgresUsername string
	PostgresPassword string
	PostgresDBName   string
}

// NewDefaultPostgresOpts returns a new PostgresOpts with default values
func NewDefaultPostgresOpts() *PostgresOpts {
	return &PostgresOpts{
		PostgresHost:     "localhost",
		PostgresPort:     5432,
		PostgresUsername: "postgres",
		PostgresPassword: "password",
		PostgresDBName:   "postgres",
	}
}

// GetPostgresOpts parses the cobra.Command and returns the PostgresOpts.
func GetPostgresOpts(c *cli.Context) *PostgresOpts {
	return &PostgresOpts{
		PostgresHost:     c.String(postgresHost),
		PostgresPort:     c.Int(postgresPort),
		PostgresUsername: c.String(postgresUsername),
		PostgresPassword: c.String(postgresPassword),
		PostgresDBName:   c.String(postgresDBName),
	}
}

func (p *PostgresOpts) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.PostgresHost,
		p.PostgresPort,
		p.PostgresUsername,
		p.PostgresPassword,
		p.PostgresDBName,
	)
}

func (p *PostgresOpts) ConnectionStringWithoutDBName() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		p.PostgresHost,
		p.PostgresPort,
		p.PostgresUsername,
		p.PostgresPassword,
	)
}

// AddPostgresFlags adds the Postgres-specific command line arguments to the cobra.Command.
func AddPostgresFlags(flags *[]cli.Flag) {
	defaultOpts := NewDefaultPostgresOpts()
	*flags = append(*flags, &cli.StringFlag{
		Name:    postgresHost,
		Value:   defaultOpts.PostgresHost,
		Usage:   "PostgreSQL `host`",
		EnvVars: []string{"PSQL_HOST"},
	})
	*flags = append(*flags, &cli.IntFlag{
		Name:    postgresPort,
		Value:   defaultOpts.PostgresPort,
		Usage:   "PostgreSQL `port`",
		EnvVars: []string{"PSQL_PORT"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    postgresUsername,
		Value:   defaultOpts.PostgresUsername,
		Usage:   "PostgreSQL `username`",
		EnvVars: []string{"PSQL_USERNAME"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    postgresPassword,
		Value:   defaultOpts.PostgresPassword,
		Usage:   "PostgreSQL `password`",
		EnvVars: []string{"PSQL_PASSWORD"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    postgresDBName,
		Value:   defaultOpts.PostgresDBName,
		Usage:   "PostgreSQL `dbname`",
		EnvVars: []string{"PSQL_DBNAME"},
	})
}
