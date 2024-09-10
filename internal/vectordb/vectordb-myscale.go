package VectorDB

import (
	Define "github.com/soulteary/dify-easy/define"
	CustomConfig "github.com/soulteary/dify-easy/internal/custom-config"
)

type Myscale struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Profiles      []string `yaml:"profiles"`
	Restart       string   `yaml:"restart"`
	Tty           bool     `yaml:"tty"`
	Volumes       []string `yaml:"volumes"`
	Ports         []string `yaml:"ports"`
}

func CreateMyscale() Myscale {
	return Myscale{
		ContainerName: "myscale",
		Image:         CustomConfig.GetImage(Define.DOCKER_SERVICE_VDB_MYSCALE),
		Profiles:      []string{"myscale"},
		Restart:       "always",
		Tty:           true,
		Volumes: []string{
			"./volumes/myscale/data:/var/lib/clickhouse",
			"./volumes/myscale/log:/var/log/clickhouse-server",
			"./volumes/myscale/config/users.d/custom_users_config.xml:/etc/clickhouse-server/users.d/custom_users_config.xml",
		},
		Ports: []string{
			"${MYSCALE_PORT:-8123}:${MYSCALE_PORT:-8123}",
		},
	}
}
