package DifyCore

import CustomConfig "github.com/soulteary/dify-easy/custom-config"

type Worker struct {
	Image   string `yaml:"image"`
	Restart string `yaml:"restart"`
	// Environment XSharedEnv `yaml:"environment"`
	Environment map[string]any `yaml:"environment"`
	DependsOn   []string       `yaml:"depends_on"`
	Volumes     []string       `yaml:"volumes"`
	Networks    []string       `yaml:"networks"`
}

func CreateDifyWorker() Worker {
	return Worker{
		Image:   CustomConfig.GetImage(CustomConfig.DOCKER_IMAGE_TYPE_DIFY_WORKER),
		Restart: "always",
		Environment: map[string]any{
			"<<":   "*shared-api-worker-env",
			"MODE": "worker",
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
