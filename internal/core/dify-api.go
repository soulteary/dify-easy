package DifyCore

import (
	Define "github.com/soulteary/dify-easy/define"
	CustomConfig "github.com/soulteary/dify-easy/internal/custom-config"
)

func CreateDifyAPI() Define.DifyAPI {
	return Define.DifyAPI{
		Image:   CustomConfig.GetImage(Define.DOCKER_SERVICE_DIFY_API),
		Restart: "always",
		Environment: Define.DifyEnvirontment{
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
