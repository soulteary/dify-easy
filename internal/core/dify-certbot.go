package DifyCore

type Certbot struct {
	Image       string   `yaml:"image"`
	Profiles    []string `yaml:"profiles"`
	Volumes     []string `yaml:"volumes"`
	Environment []string `yaml:"environment"`
	Entrypoint  []string `yaml:"entrypoint"`
	Command     []string `yaml:"command"`
}

func CreateDifyCertbot() Certbot {
	return Certbot{
		Image:    "certbot/certbot",
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
		Entrypoint: []string{
			"/docker-entrypoint.sh",
		},
		Command: []string{
			"tail",
			"-f",
			"/dev/null",
		},
	}
}
