package DifyCore

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
)

type Unstructured struct {
	Image    string   `yaml:"image"`
	Profiles []string `yaml:"profiles"`
	Restart  string   `yaml:"restart"`
	Volumes  []string `yaml:"volumes"`
}

func CreateUnstructured() Unstructured {
	return Unstructured{
		Image:    CustomConfig.GetImage(Define.DOCKER_SERVICE_DIFY_UNSTRUCTURED),
		Profiles: []string{"unstructured"},
		Restart:  "always",
		Volumes: []string{
			"./volumes/unstructured:/app/data",
		},
	}
}
