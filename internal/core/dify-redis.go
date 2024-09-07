package DifyCore

type Redis struct {
	Image       string      `yaml:"image"`
	Restart     string      `yaml:"restart"`
	Volumes     []string    `yaml:"volumes"`
	Command     string      `yaml:"command"`
	Healthcheck HealthCheck `yaml:"healthcheck"`
}

func CreateDifyRedis() Redis {
	return Redis{
		Image:   "redis:6-alpine",
		Restart: "always",
		Volumes: []string{
			"./volumes/redis/data:/data",
		},
		Command: "redis-server --requirepass ${REDIS_PASSWORD:-difyai123456}",
		Healthcheck: HealthCheck{
			Test: []string{"CMD", "redis-cli", "ping"},
		},
	}
}
