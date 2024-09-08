package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Fn "github.com/soulteary/dify-easy/fn"
)

type Nginx struct {
	Image       string           `yaml:"image"`
	Restart     string           `yaml:"restart"`
	Volumes     []string         `yaml:"volumes"`
	Entrypoint  string           `yaml:"entrypoint"`
	Environment NginxEnvironment `yaml:"environment"`
	DependsOn   []string         `yaml:"depends_on"`
	Ports       []string         `yaml:"ports"`
}

type NginxEnvironment struct {
	NGINXSERVERNAME             string `yaml:"NGINX_SERVER_NAME"`
	NGINXHTTPSENABLED           string `yaml:"NGINX_HTTPS_ENABLED"`
	NGINXSSLPORT                string `yaml:"NGINX_SSL_PORT"`
	NGINXPORT                   string `yaml:"NGINX_PORT"`
	NGINXSSLCERTFILENAME        string `yaml:"NGINX_SSL_CERT_FILENAME"`
	NGINXSSLCERTKEYFILENAME     string `yaml:"NGINX_SSL_CERT_KEY_FILENAME"`
	NGINXSSLPROTOCOLS           string `yaml:"NGINX_SSL_PROTOCOLS"`
	NGINXWORKERPROCESSES        string `yaml:"NGINX_WORKER_PROCESSES"`
	NGINXCLIENTMAXBODYSIZE      string `yaml:"NGINX_CLIENT_MAX_BODY_SIZE"`
	NGINXKEEPALIVETIMEOUT       string `yaml:"NGINX_KEEPALIVE_TIMEOUT"`
	NGINXPROXYREADTIMEOUT       string `yaml:"NGINX_PROXY_READ_TIMEOUT"`
	NGINXPROXYSENDTIMEOUT       string `yaml:"NGINX_PROXY_SEND_TIMEOUT"`
	NGINXENABLECERTBOTCHALLENGE string `yaml:"NGINX_ENABLE_CERTBOT_CHALLENGE"`
	CERTBOTDOMAIN               string `yaml:"CERTBOT_DOMAIN"`
}

func CreateDifyNginx() Nginx {
	config := Nginx{
		Image:   CustomConfig.GetImage(CustomConfig.DOCKER_SERVICE_DIFY_NGINX),
		Restart: "always",
		Volumes: []string{
			"./nginx/nginx.conf.template:/etc/nginx/nginx.conf.template",
			"./nginx/proxy.conf.template:/etc/nginx/proxy.conf.template",
			"./nginx/https.conf.template:/etc/nginx/https.conf.template",
			"./nginx/conf.d:/etc/nginx/conf.d",
			"./nginx/docker-entrypoint.sh:/docker-entrypoint-mount.sh",
			"./nginx/ssl:/etc/ssl",
			"./volumes/certbot/conf/live:/etc/letsencrypt/live",
			"./volumes/certbot/conf:/etc/letsencrypt",
			"./volumes/certbot/www:/var/www/html",
		},
		Environment: NginxEnvironment{
			NGINXSERVERNAME:             "${NGINX_SERVER_NAME:-_}",
			NGINXHTTPSENABLED:           "${NGINX_HTTPS_ENABLED:-false}",
			NGINXSSLPORT:                "${NGINX_SSL_PORT:-443}",
			NGINXPORT:                   "${NGINX_PORT:-80}",
			NGINXSSLCERTFILENAME:        "${NGINX_SSL_CERT_FILENAME:-dify.crt}",
			NGINXSSLCERTKEYFILENAME:     "${NGINX_SSL_CERT_KEY_FILENAME:-dify.key}",
			NGINXSSLPROTOCOLS:           "${NGINX_SSL_PROTOCOLS:-TLSv1.1 TLSv1.2 TLSv1.3}",
			NGINXWORKERPROCESSES:        "${NGINX_WORKER_PROCESSES:-auto}",
			NGINXCLIENTMAXBODYSIZE:      "${NGINX_CLIENT_MAX_BODY_SIZE:-15M}",
			NGINXKEEPALIVETIMEOUT:       "${NGINX_KEEPALIVE_TIMEOUT:-65}",
			NGINXPROXYREADTIMEOUT:       "${NGINX_PROXY_READ_TIMEOUT:-3600s}",
			NGINXPROXYSENDTIMEOUT:       "${NGINX_PROXY_SEND_TIMEOUT:-3600s}",
			NGINXENABLECERTBOTCHALLENGE: "${NGINX_ENABLE_CERTBOT_CHALLENGE:-false}",
			CERTBOTDOMAIN:               "${CERTBOT_DOMAIN:-}",
		},
		DependsOn: []string{
			"api",
			"web",
		},
		Ports: []string{
			`"${EXPOSE_NGINX_PORT:-80}:${NGINX_PORT:-80}"`,
			`"${EXPOSE_NGINX_SSL_PORT:-443}:${NGINX_SSL_PORT:-443}"`,
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
