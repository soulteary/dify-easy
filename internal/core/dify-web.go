package DifyCore

import CustomConfig "github.com/soulteary/dify-easy/custom-config"

type Web struct {
	Image       string          `yaml:"image"`
	Restart     string          `yaml:"restart"`
	Environment WebEnvirontment `yaml:"environment"`
}

type WebEnvirontment struct {
	CONSOLEAPIURL         string `yaml:"CONSOLE_API_URL"`
	APPAPIURL             string `yaml:"APP_API_URL"`
	SENTRYDSN             string `yaml:"SENTRY_DSN"`
	NEXTTELEMETRYDISABLED string `yaml:"NEXT_TELEMETRY_DISABLED"`
}

func CreateDifyWeb() Web {
	return Web{
		Image:   CustomConfig.GetImage(CustomConfig.DOCKER_IMAGE_TYPE_DIFY_WEB),
		Restart: "always",
		Environment: WebEnvirontment{
			CONSOLEAPIURL:         "${CONSOLE_API_URL:-}",
			APPAPIURL:             "${APP_API_URL:-}",
			SENTRYDSN:             "${WEB_SENTRY_DSN:-}",
			NEXTTELEMETRYDISABLED: "${NEXT_TELEMETRY_DISABLED:-0}",
		},
	}
}
