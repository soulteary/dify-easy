package CustomConfig

const (
	DEFAULT_DOCKER_IMAGE_DIFY_API         = "langgenius/dify-api"
	DEFAULT_DOCKER_IMAGE_VERSION_DIFY_API = "0.7.3"

	DEFAULT_DOCKER_IMAGE_DIFY_WORKER         = "langgenius/dify-api"
	DEFAULT_DOCKER_IMAGE_VERSION_DIFY_WORKER = "0.7.3"

	DEFAULT_DOCKER_IMAGE_DIFY_WEB         = "langgenius/dify-web"
	DEFAULT_DOCKER_IMAGE_VERSION_DIFY_WEB = "0.7.3"

	DEFAULT_DOCKER_IMAGE_DIFY_REDIS         = "redis"
	DEFAULT_DOCKER_IMAGE_VERSION_DIFY_REDIS = "6-alpine"

	DEFAULT_DOCKER_IMAGE_DIFY_POSTGRES         = "postgres"
	DEFAULT_DOCKER_IMAGE_VERSION_DIFY_POSTGRES = "15-alpine"

	DEFAULT_DOCKER_IMAGE_DIFY_NGINX         = "nginx"
	DEFAULT_DOCKER_IMAGE_VERSION_DIFY_NGINX = "latest"

	DEFAULT_DOCKER_IMAGE_DIFY_SANDBOX         = "langgenius/dify-sandbox"
	DEFAULT_DOCKER_IMAGE_VERSION_DIFY_SANDBOX = "0.2.6"

	DEFAULT_DOCKER_IMAGE_DIFY_SSRF_PROXY         = "ubuntu/squid"
	DEFAULT_DOCKER_IMAGE_VERSION_DIFY_SSRF_PROXY = "latest"

	DEFAULT_DOCKER_IMAGE_DIFY_CERTBOT         = "certbot/certbot"
	DEFAULT_DOCKER_IMAGE_VERSION_DIFY_CERTBOT = ""
)

const (
	DOCKER_IMAGE_TYPE_DIFY_API = iota
	DOCKER_IMAGE_TYPE_DIFY_WORKER
	DOCKER_IMAGE_TYPE_DIFY_WEB
	DOCKER_IMAGE_TYPE_DIFY_REDIS
	DOCKER_IMAGE_TYPE_DIFY_POSTGRES
	DOCKER_IMAGE_TYPE_DIFY_NGINX
	DOCKER_IMAGE_TYPE_DIFY_SANDBOX
	DOCKER_IMAGE_TYPE_DIFY_SSRF_PROXY
	DOCKER_IMAGE_TYPE_DIFY_CERTBOT
)

func GetImage(imageType int) string {
	switch imageType {
	case DOCKER_IMAGE_TYPE_DIFY_API:
		return DEFAULT_DOCKER_IMAGE_DIFY_API + ":" + DEFAULT_DOCKER_IMAGE_VERSION_DIFY_API
	case DOCKER_IMAGE_TYPE_DIFY_WORKER:
		return DEFAULT_DOCKER_IMAGE_DIFY_WORKER + ":" + DEFAULT_DOCKER_IMAGE_VERSION_DIFY_WORKER
	case DOCKER_IMAGE_TYPE_DIFY_WEB:
		return DEFAULT_DOCKER_IMAGE_DIFY_WEB + ":" + DEFAULT_DOCKER_IMAGE_VERSION_DIFY_WEB
	case DOCKER_IMAGE_TYPE_DIFY_REDIS:
		return DEFAULT_DOCKER_IMAGE_DIFY_REDIS + ":" + DEFAULT_DOCKER_IMAGE_VERSION_DIFY_REDIS
	case DOCKER_IMAGE_TYPE_DIFY_POSTGRES:
		return DEFAULT_DOCKER_IMAGE_DIFY_POSTGRES + ":" + DEFAULT_DOCKER_IMAGE_VERSION_DIFY_POSTGRES
	case DOCKER_IMAGE_TYPE_DIFY_NGINX:
		return DEFAULT_DOCKER_IMAGE_DIFY_NGINX + ":" + DEFAULT_DOCKER_IMAGE_VERSION_DIFY_NGINX
	case DOCKER_IMAGE_TYPE_DIFY_SANDBOX:
		return DEFAULT_DOCKER_IMAGE_DIFY_SANDBOX + ":" + DEFAULT_DOCKER_IMAGE_VERSION_DIFY_SANDBOX
	case DOCKER_IMAGE_TYPE_DIFY_SSRF_PROXY:
		return DEFAULT_DOCKER_IMAGE_DIFY_SSRF_PROXY + ":" + DEFAULT_DOCKER_IMAGE_VERSION_DIFY_SSRF_PROXY
	case DOCKER_IMAGE_TYPE_DIFY_CERTBOT:
		return DEFAULT_DOCKER_IMAGE_DIFY_CERTBOT + ":" + DEFAULT_DOCKER_IMAGE_VERSION_DIFY_CERTBOT
	}

	return "not-found"
}
