package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
)

func CreateDifyWeb() Define.DifyWeb {
	return Define.DifyWeb{
		Image:   CustomConfig.GetImage(Define.DOCKER_SERVICE_DIFY_WEB),
		Restart: "always",
		Environment: Define.DifyWebEnvirontment{
			CONSOLE_API_URL:         "${CONSOLE_API_URL:-}",
			APP_API_URL:             "${APP_API_URL:-}",
			SENTRY_DSN:              "${WEB_SENTRY_DSN:-}",
			NEXT_TELEMETRY_DISABLED: "${NEXT_TELEMETRY_DISABLED:-0}",
		},
	}
}
