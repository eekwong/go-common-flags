package flag

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const (
	mongoHost     = "mongo-host"
	mongoPort     = "mongo-port"
	mongoUsername = "mongo-username"
	mongoPassword = "mongo-password"
)

type MongoOpts struct {
	MongoHost     string
	MongoPort     int
	MongoUsername string
	MongoPassword string
}

func NewDefaultMongoOpts() *MongoOpts {
	return &MongoOpts{
		MongoHost:     "localhost",
		MongoPort:     27017,
		MongoUsername: "admin",
		MongoPassword: "password",
	}
}

func GetMongoOpts(c *cli.Context) *MongoOpts {
	return &MongoOpts{
		MongoHost:     c.String(mongoHost),
		MongoPort:     c.Int(mongoPort),
		MongoUsername: c.String(mongoUsername),
		MongoPassword: c.String(mongoPassword),
	}
}

func (m *MongoOpts) GetURI() string {
	if m.MongoUsername != "" && m.MongoPassword != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%d/",
			m.MongoUsername,
			m.MongoPassword,
			m.MongoHost,
			m.MongoPort,
		)
	}
	return fmt.Sprintf("mongodb://%s:%d/",
		m.MongoHost,
		m.MongoPort,
	)
}

func AddMongoFlags(flags *[]cli.Flag) {
	defaultOpts := NewDefaultMongoOpts()
	*flags = append(*flags, &cli.StringFlag{
		Name:    mongoHost,
		Value:   defaultOpts.MongoHost,
		Usage:   "Mongo `host`",
		EnvVars: []string{"MONGO_HOST"},
	})
	*flags = append(*flags, &cli.IntFlag{
		Name:    mongoPort,
		Value:   defaultOpts.MongoPort,
		Usage:   "Mongo `port`",
		EnvVars: []string{"MONGO_PORT"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    mongoUsername,
		Value:   defaultOpts.MongoUsername,
		Usage:   "Mongo `username`",
		EnvVars: []string{"MONGO_USERNAME"},
	})
	*flags = append(*flags, &cli.StringFlag{
		Name:    mongoPassword,
		Value:   defaultOpts.MongoPassword,
		Usage:   "Mongo `password`",
		EnvVars: []string{"MONGO_PASSWORD"},
	})
}
