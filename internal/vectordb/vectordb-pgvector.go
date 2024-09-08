package VectorDB

import (
	Define "github.com/soulteary/dify-easy/define"
	Fn "github.com/soulteary/dify-easy/fn"
)

type Pgvector struct {
	Image       string              `yaml:"image"`
	Profiles    []string            `yaml:"profiles"`
	Restart     string              `yaml:"restart"`
	Environment PgvectorEnvironment `yaml:"environment"`
	Volumes     []string            `yaml:"volumes"`
	Healthcheck Define.HealthCheck  `yaml:"healthcheck"`
}

type PgvectorEnvironment struct {
	PGUSER            string `yaml:"PGUSER"`
	POSTGRES_PASSWORD string `yaml:"POSTGRES_PASSWORD"`
	POSTGRES_DB       string `yaml:"POSTGRES_DB"`
	PGDATA            string `yaml:"PGDATA"`
}

func CreatePgvector() Pgvector {
	config := Pgvector{
		Image:    "pgvector/pgvector:pg16",
		Profiles: []string{"pgvector"},
		Restart:  "always",
		Environment: PgvectorEnvironment{
			PGUSER:            "${PGVECTOR_PGUSER:-postgres}",
			POSTGRES_PASSWORD: "${PGVECTOR_POSTGRES_PASSWORD:-difyai123456}",
			POSTGRES_DB:       "${PGVECTOR_POSTGRES_DB:-dify}",
			PGDATA:            "${PGVECTOR_PGDATA:-/var/lib/postgresql/data/pgdata}",
		},
		Volumes: []string{
			"./volumes/pgvector/data:/var/lib/postgresql/data",
		},
	}

	healthCheckCmd, err := Fn.ConvertArrToCommand([]string{"CMD", "pg_isready"})
	if err == nil {
		config.Healthcheck = Define.HealthCheck{
			Test: healthCheckCmd,
		}
		config.Healthcheck.Interval = "1s"
		config.Healthcheck.Timeout = "3s"
		config.Healthcheck.Retries = 30
	}

	return config
}
