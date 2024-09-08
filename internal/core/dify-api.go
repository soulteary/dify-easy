package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
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
