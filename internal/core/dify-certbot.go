package DifyCore

import (
	Define "github.com/soulteary/dify-easy/define"
	CustomConfig "github.com/soulteary/dify-easy/internal/custom-config"
	Fn "github.com/soulteary/dify-easy/internal/fn"
)

type Certbot struct {
	Image       string   `yaml:"image"`
	Profiles    []string `yaml:"profiles"`
	Volumes     []string `yaml:"volumes"`
	Environment []string `yaml:"environment"`
	Entrypoint  string   `yaml:"entrypoint"`
	Command     string   `yaml:"command"`
}

func CreateDifyCertbot() Certbot {
	config := Certbot{
		Image:    CustomConfig.GetImage(Define.DOCKER_SERVICE_DIFY_CERTBOT),
		Profiles: []string{"certbot"},
		Volumes: []string{
			"./volumes/certbot/conf:/etc/letsencrypt",
			"./volumes/certbot/www:/var/www/html",
			"./volumes/certbot/logs:/var/log/letsencrypt",
			"./volumes/certbot/conf/live:/etc/letsencrypt/live",
			"./certbot/update-cert.template.txt:/update-cert.template.txt",
			"./certbot/docker-entrypoint.sh:/docker-entrypoint.sh",
		},
		Environment: []string{
			"CERTBOT_EMAIL=${CERTBOT_EMAIL}",
			"CERTBOT_DOMAIN=${CERTBOT_DOMAIN}",
			"CERTBOT_OPTIONS=${CERTBOT_OPTIONS:-}",
		},
	}

	entrypointCommand, err := Fn.ConvertArrToCommand([]string{
		"/docker-entrypoint.sh",
	})

	if err == nil {
		config.Entrypoint = entrypointCommand
	}

	command, err := Fn.ConvertArrToCommand([]string{
		"tail",
		"-f",
		"/dev/null",
	})

	if err == nil {
		config.Command = command
	}

	return config
}
