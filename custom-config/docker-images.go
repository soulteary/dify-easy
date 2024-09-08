package CustomConfig

import (
	Define "github.com/soulteary/dify-easy/define"
)

type DockerImage struct {
	Image   string `json:"image"`
	Version string `json:"version"`
}

type DockerImages map[int]DockerImage

var DefaultDockerImages = func() DockerImages {
	return DockerImages{
		Define.DOCKER_SERVICE_DIFY_API: {
			Image:   Define.DEFAULT_DOCKER_IMAGE_DIFY_API,
			Version: Define.DEFAULT_DOCKER_IMAGE_VERSION_DIFY_API,
		},
		Define.DOCKER_SERVICE_DIFY_WORKER: {
			Image:   Define.DEFAULT_DOCKER_IMAGE_DIFY_WORKER,
			Version: Define.DEFAULT_DOCKER_IMAGE_VERSION_DIFY_WORKER,
		},
		Define.DOCKER_SERVICE_DIFY_WEB: {
			Image:   Define.DEFAULT_DOCKER_IMAGE_DIFY_WEB,
			Version: Define.DEFAULT_DOCKER_IMAGE_VERSION_DIFY_WEB,
		},
		Define.DOCKER_SERVICE_DIFY_REDIS: {
			Image:   Define.DEFAULT_DOCKER_IMAGE_DIFY_REDIS,
			Version: Define.DEFAULT_DOCKER_IMAGE_VERSION_DIFY_REDIS,
		},
		Define.DOCKER_SERVICE_DIFY_POSTGRES: {
			Image:   Define.DEFAULT_DOCKER_IMAGE_DIFY_POSTGRES,
			Version: Define.DEFAULT_DOCKER_IMAGE_VERSION_DIFY_POSTGRES,
		},
		Define.DOCKER_SERVICE_DIFY_NGINX: {
			Image:   Define.DEFAULT_DOCKER_IMAGE_DIFY_NGINX,
			Version: Define.DEFAULT_DOCKER_IMAGE_VERSION_DIFY_NGINX,
		},
		Define.DOCKER_SERVICE_DIFY_SANDBOX: {
			Image:   Define.DEFAULT_DOCKER_IMAGE_DIFY_SANDBOX,
			Version: Define.DEFAULT_DOCKER_IMAGE_VERSION_DIFY_SANDBOX,
		},
		Define.DOCKER_SERVICE_DIFY_SSRF_PROXY: {
			Image:   Define.DEFAULT_DOCKER_IMAGE_DIFY_SSRF_PROXY,
			Version: Define.DEFAULT_DOCKER_IMAGE_VERSION_DIFY_SSRF_PROXY,
		},
		Define.DOCKER_SERVICE_DIFY_CERTBOT: {
			Image:   Define.DEFAULT_DOCKER_IMAGE_DIFY_CERTBOT,
			Version: Define.DEFAULT_DOCKER_IMAGE_VERSION_DIFY_CERTBOT,
		},
		Define.DOCKER_SERVICE_DIFY_UNSTRUCTURED: {
			Image:   Define.DEFAULT_DOCKER_IMAGE_DIFY_UNSTRUCTURED,
			Version: Define.DEFAULT_DOCKER_IMAGE_DIFY_UNSTRUCTURED_VERSION,
		},
	}
}()

func GetImage(imageType int) string {
	info, ok := DefaultDockerImages[imageType]
	if ok {
		if info.Version == "" {
			return info.Image
		}
		return info.Image + ":" + info.Version
	}
	return "not-found"
}
