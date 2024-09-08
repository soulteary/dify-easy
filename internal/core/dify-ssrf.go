package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
	Fn "github.com/soulteary/dify-easy/fn"
)

type SsrfProxy struct {
	Image       string               `yaml:"image"`
	Restart     string               `yaml:"restart"`
	Volumes     []string             `yaml:"volumes"`
	Entrypoint  string               `yaml:"entrypoint"`
	Environment SsrfProxyEnvironment `yaml:"environment"`
	Networks    []string             `yaml:"networks"`
}

type SsrfProxyEnvironment struct {
	HTTPPORT         string `yaml:"HTTP_PORT"`
	COREDUMPDIR      string `yaml:"COREDUMP_DIR"`
	REVERSEPROXYPORT string `yaml:"REVERSE_PROXY_PORT"`
	SANDBOXHOST      string `yaml:"SANDBOX_HOST"`
	SANDBOXPORT      string `yaml:"SANDBOX_PORT"`
}

func CreateDifySsrfProxy() SsrfProxy {

	config := SsrfProxy{
		Image:   CustomConfig.GetImage(Define.DOCKER_SERVICE_DIFY_SSRF_PROXY),
		Restart: "always",
		Volumes: []string{
			"./ssrf_proxy/squid.conf.template:/etc/squid/squid.conf.template",
			"./ssrf_proxy/docker-entrypoint.sh:/docker-entrypoint-mount.sh",
		},
		Environment: SsrfProxyEnvironment{
			HTTPPORT:         "${SSRF_HTTP_PORT:-3128}",
			COREDUMPDIR:      "${SSRF_COREDUMP_DIR:-/var/spool/squid}",
			REVERSEPROXYPORT: "${SSRF_REVERSE_PROXY_PORT:-8194}",
			SANDBOXHOST:      "${SSRF_SANDBOX_HOST:-sandbox}",
			SANDBOXPORT:      "${SANDBOX_PORT:-8194}",
		},
		Networks: []string{
			"ssrf_proxy_network",
			"default",
		},
	}

	entrypointCommand, err := Fn.ConvertArrToCommand([]string{
		"sh",
		"-c",
		"cp /docker-entrypoint-mount.sh /docker-entrypoint.sh && sed -i 's/\\r$$//' /docker-entrypoint.sh && chmod +x /docker-entrypoint.sh && /docker-entrypoint.sh",
	})

	if err == nil {
		config.Entrypoint = entrypointCommand
	}

	return config
}
