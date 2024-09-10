package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
	Fn "github.com/soulteary/dify-easy/fn"
)

type Sandbox struct {
	Image       string             `yaml:"image"`
	Restart     string             `yaml:"restart"`
	Environment SandboxEnvironment `yaml:"environment"`
	Volumes     []string           `yaml:"volumes"`
	Networks    []string           `yaml:"networks"`
	Healthcheck Define.HealthCheck `yaml:"healthcheck"`
}

type SandboxEnvironment struct {
	APIKEY        string `yaml:"API_KEY"`
	GINMODE       string `yaml:"GIN_MODE"`
	WORKERTIMEOUT string `yaml:"WORKER_TIMEOUT"`
	ENABLENETWORK string `yaml:"ENABLE_NETWORK"`
	HTTPPROXY     string `yaml:"HTTP_PROXY"`
	HTTPSPROXY    string `yaml:"HTTPS_PROXY"`
	SANDBOXPORT   string `yaml:"SANDBOX_PORT"`
}

func CreateDifySandbox() Sandbox {
	config := Sandbox{
		Image:   CustomConfig.GetImage(Define.DOCKER_SERVICE_DIFY_SANDBOX),
		Restart: "always",
		Environment: SandboxEnvironment{
			APIKEY:        "${SANDBOX_API_KEY:-dify-sandbox}",
			GINMODE:       "${SANDBOX_GIN_MODE:-release}",
			WORKERTIMEOUT: "${SANDBOX_WORKER_TIMEOUT:-15}",
			ENABLENETWORK: "${SANDBOX_ENABLE_NETWORK:-true}",
			HTTPPROXY:     "${SANDBOX_HTTP_PROXY:-http://ssrf_proxy:3128}",
			HTTPSPROXY:    "${SANDBOX_HTTPS_PROXY:-http://ssrf_proxy:3128}",
			SANDBOXPORT:   "${SANDBOX_PORT:-8194}",
		},
		Volumes: []string{
			"./volumes/sandbox/dependencies:/dependencies",
		},
		Networks: []string{
			"ssrf_proxy_network",
		},
	}

	healthCheckCmd, err := Fn.ConvertArrToCommand([]string{"CMD", "curl", "-f", "http://localhost:8194/health"})
	if err == nil {
		config.Healthcheck = Define.HealthCheck{
			Test: healthCheckCmd,
		}
	}
	return config
}
