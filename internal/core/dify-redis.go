package DifyCore

import (
	Define "github.com/soulteary/dify-easy/define"
	CustomConfig "github.com/soulteary/dify-easy/internal/custom-config"
	Fn "github.com/soulteary/dify-easy/internal/fn"
)

type Redis struct {
	Image       string             `yaml:"image"`
	Restart     string             `yaml:"restart"`
	Volumes     []string           `yaml:"volumes"`
	Command     string             `yaml:"command"`
	Healthcheck Define.HealthCheck `yaml:"healthcheck"`
}

func CreateDifyRedis() Redis {
	config := Redis{
		Image:   CustomConfig.GetImage(Define.DOCKER_SERVICE_DIFY_REDIS),
		Restart: "always",
		Volumes: []string{
			"./volumes/redis/data:/data",
		},
		Command: "redis-server --requirepass ${REDIS_PASSWORD:-difyai123456}",
	}

	healthCheckCmd, err := Fn.ConvertArrToCommand([]string{"CMD", "redis-cli", "ping"})
	if err == nil {
		config.Healthcheck = Define.HealthCheck{
			Test: healthCheckCmd,
		}
	}
	return config
}
