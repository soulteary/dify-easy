package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
)

type Web struct {
	Image       string          `yaml:"image"`
	Restart     string          `yaml:"restart"`
	Environment WebEnvirontment `yaml:"environment"`
}

type WebEnvirontment struct {
	CONSOLE_API_URL         string `yaml:"CONSOLE_API_URL"`
	APP_API_URL             string `yaml:"APP_API_URL"`
	SENTRY_DSN              string `yaml:"SENTRY_DSN"`
	NEXT_TELEMETRY_DISABLED string `yaml:"NEXT_TELEMETRY_DISABLED"`
}

func CreateDifyWeb() Web {
	return Web{
		Image:   CustomConfig.GetImage(Define.DOCKER_SERVICE_DIFY_WEB),
		Restart: "always",
		Environment: WebEnvirontment{
			CONSOLE_API_URL:         "${CONSOLE_API_URL:-}",
			APP_API_URL:             "${APP_API_URL:-}",
			SENTRY_DSN:              "${WEB_SENTRY_DSN:-}",
			NEXT_TELEMETRY_DISABLED: "${NEXT_TELEMETRY_DISABLED:-0}",
		},
	}
}
