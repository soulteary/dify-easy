package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
)

type API struct {
	Image   string `yaml:"image"`
	Restart string `yaml:"restart"`
	// Environment XSharedEnv `yaml:"environment"`
	Environment map[string]any `yaml:"environment"`
	DependsOn   []string       `yaml:"depends_on"`
	Volumes     []string       `yaml:"volumes"`
	Networks    []string       `yaml:"networks"`
}

func CreateDifyAPI() API {
	return API{
		Image:   CustomConfig.GetImage(CustomConfig.DOCKER_SERVICE_DIFY_API),
		Restart: "always",
		Environment: map[string]any{
			"<<":   "*shared-api-worker-env",
			"MODE": "api",
		},
		DependsOn: []string{
			"db",
			"redis",
		},
		Volumes: []string{
			"./volumes/app/storage:/app/api/storage",
		},
		Networks: []string{
			"ssrf_proxy_network",
			"default",
		},
	}
}
