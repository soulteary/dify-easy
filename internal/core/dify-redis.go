package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Fn "github.com/soulteary/dify-easy/fn"
)

type Redis struct {
	Image       string      `yaml:"image"`
	Restart     string      `yaml:"restart"`
	Volumes     []string    `yaml:"volumes"`
	Command     string      `yaml:"command"`
	Healthcheck HealthCheck `yaml:"healthcheck"`
}

func CreateDifyRedis() Redis {
	config := Redis{
		Image:   CustomConfig.GetImage(CustomConfig.DOCKER_IMAGE_TYPE_DIFY_REDIS),
		Restart: "always",
		Volumes: []string{
			"./volumes/redis/data:/data",
		},
		Command: "redis-server --requirepass ${REDIS_PASSWORD:-difyai123456}",
	}

	healthCheckCmd, err := Fn.ConvertArrToCommand([]string{"CMD", "redis-cli", "ping"})
	if err == nil {
		config.Healthcheck = HealthCheck{
			Test: healthCheckCmd,
		}
	}
	return config
}
