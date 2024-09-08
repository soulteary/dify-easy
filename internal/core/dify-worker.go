package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
)

type Worker struct {
	Image   string `yaml:"image"`
	Restart string `yaml:"restart"`
	// Environment XSharedEnv `yaml:"environment"`
	Environment Define.DifyEnvirontment `yaml:"environment"`
	DependsOn   []string                `yaml:"depends_on"`
	Volumes     []string                `yaml:"volumes"`
	Networks    []string                `yaml:"networks"`
}

func CreateDifyWorker() Worker {
	return Worker{
		Image:   CustomConfig.GetImage(Define.DOCKER_SERVICE_DIFY_WORKER),
		Restart: "always",
		Environment: Define.DifyEnvirontment{
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
