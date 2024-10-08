package DifyCore

import (
	"strings"

	Define "github.com/soulteary/dify-easy/define"
	CustomConfig "github.com/soulteary/dify-easy/internal/custom-config"
	Fn "github.com/soulteary/dify-easy/internal/fn"
)

type DB struct {
	Image       string             `yaml:"image"`
	Restart     string             `yaml:"restart"`
	Environment DBEnvironment      `yaml:"environment"`
	Command     string             `yaml:"command"`
	Volumes     []string           `yaml:"volumes"`
	Healthcheck Define.HealthCheck `yaml:"healthcheck"`
}

type DBEnvironment struct {
	PGUSER            string `yaml:"PGUSER"`
	POSTGRES_PASSWORD string `yaml:"POSTGRES_PASSWORD"`
	POSTGRES_DB       string `yaml:"POSTGRES_DB"`
	PGDATA            string `yaml:"PGDATA"`
}

func CreateDifyDB() DB {
	config := DB{
		Image:   CustomConfig.GetImage(Define.DOCKER_SERVICE_DIFY_POSTGRES),
		Restart: "always",
		Environment: DBEnvironment{
			PGUSER:            "${PGUSER:-postgres}",
			POSTGRES_PASSWORD: "${POSTGRES_PASSWORD:-difyai123456}",
			POSTGRES_DB:       "${POSTGRES_DB:-dify}",
			PGDATA:            "${PGDATA:-/var/lib/postgresql/data/pgdata}",
		},
		Command: strings.Join([]string{
			">",
			"postgres -c 'max_connections=${POSTGRES_MAX_CONNECTIONS:-100}'",
			"         -c 'shared_buffers=${POSTGRES_SHARED_BUFFERS:-128MB}'",
			"         -c 'work_mem=${POSTGRES_WORK_MEM:-4MB}'",
			"         -c 'maintenance_work_mem=${POSTGRES_MAINTENANCE_WORK_MEM:-64MB}'",
			"         -c 'effective_cache_size=${POSTGRES_EFFECTIVE_CACHE_SIZE:-4096MB}'",
		}, "\n"),
		Volumes: []string{
			"./volumes/db/data:/var/lib/postgresql/data",
		},
	}

	healthCheckCmd, err := Fn.ConvertArrToCommand([]string{"CMD", "pg_isready"})
	if err == nil {
		config.Healthcheck = Define.HealthCheck{
			Test:     healthCheckCmd,
			Interval: "1s",
			Timeout:  "3s",
			Retries:  30,
		}
	}
	return config
}
