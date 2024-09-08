package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
)

func CreateDifyWorker() Define.DifyWorker {
	return Define.DifyWorker{
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
