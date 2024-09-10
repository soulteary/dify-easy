package VectorDB

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
)

type Weaviate struct {
	Image       string              `yaml:"image"`
	Profiles    []string            `yaml:"profiles"`
	Restart     string              `yaml:"restart"`
	Volumes     []string            `yaml:"volumes"`
	Environment WeaviateEnvironment `yaml:"environment"`
}

type WeaviateEnvironment struct {
	PERSISTENCE_DATA_PATH                   string `yaml:"PERSISTENCE_DATA_PATH"`
	QUERY_DEFAULTS_LIMIT                    string `yaml:"QUERY_DEFAULTS_LIMIT"`
	AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED string `yaml:"AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED"`
	DEFAULT_VECTORIZER_MODULE               string `yaml:"DEFAULT_VECTORIZER_MODULE"`
	CLUSTER_HOSTNAME                        string `yaml:"CLUSTER_HOSTNAME"`
	AUTHENTICATION_APIKEY_ENABLED           string `yaml:"AUTHENTICATION_APIKEY_ENABLED"`
	AUTHENTICATION_APIKEY_ALLOWED_KEYS      string `yaml:"AUTHENTICATION_APIKEY_ALLOWED_KEYS"`
	AUTHENTICATION_APIKEY_USERS             string `yaml:"AUTHENTICATION_APIKEY_USERS"`
	AUTHORIZATION_ADMINLIST_ENABLED         string `yaml:"AUTHORIZATION_ADMINLIST_ENABLED"`
	AUTHORIZATION_ADMINLIST_USERS           string `yaml:"AUTHORIZATION_ADMINLIST_USERS"`
}

func CreateWeaviate() Weaviate {
	config := Weaviate{
		Image: CustomConfig.GetImage(Define.DOCKER_SERVICE_VDB_WEAVIATE),
		Profiles: []string{
			"",
			"weaviate",
		},
		Restart: "always",
		Volumes: []string{
			"./volumes/weaviate:/var/lib/weaviate",
		},
		Environment: WeaviateEnvironment{
			PERSISTENCE_DATA_PATH:                   "${WEAVIATE_PERSISTENCE_DATA_PATH:-/var/lib/weaviate}",
			QUERY_DEFAULTS_LIMIT:                    "${WEAVIATE_QUERY_DEFAULTS_LIMIT:-25}",
			AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED: "${WEAVIATE_AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED:-false}",
			DEFAULT_VECTORIZER_MODULE:               "${WEAVIATE_DEFAULT_VECTORIZER_MODULE:-none}",
			CLUSTER_HOSTNAME:                        "${WEAVIATE_CLUSTER_HOSTNAME:-node1}",
			AUTHENTICATION_APIKEY_ENABLED:           "${WEAVIATE_AUTHENTICATION_APIKEY_ENABLED:-true}",
			AUTHENTICATION_APIKEY_ALLOWED_KEYS:      "${WEAVIATE_AUTHENTICATION_APIKEY_ALLOWED_KEYS:-WVF5YThaHlkYwhGUSmCRgsX3tD5ngdN8pkih}",
			AUTHENTICATION_APIKEY_USERS:             "${WEAVIATE_AUTHENTICATION_APIKEY_USERS:-hello@dify.ai}",
			AUTHORIZATION_ADMINLIST_ENABLED:         "${WEAVIATE_AUTHORIZATION_ADMINLIST_ENABLED:-true}",
			AUTHORIZATION_ADMINLIST_USERS:           "${WEAVIATE_AUTHORIZATION_ADMINLIST_USERS:-hello@dify.ai}",
		},
	}
	return config
}
