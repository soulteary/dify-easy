package VectorDB

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
	Fn "github.com/soulteary/dify-easy/fn"
)

type Elasticsearch struct {
	Image         string                   `yaml:"image"`
	ContainerName string                   `yaml:"container_name"`
	Profiles      []string                 `yaml:"profiles"`
	Restart       string                   `yaml:"restart"`
	Volumes       []string                 `yaml:"volumes"`
	Environment   ElasticsearchEnvironment `yaml:"environment"`
	Ports         []string                 `yaml:"ports"`
	Healthcheck   Define.HealthCheck       `yaml:"healthcheck"`
}

type ElasticsearchEnvironment struct {
	ELASTIC_PASSWORD string `yaml:"ELASTIC_PASSWORD"`
	Cluster_name     string `yaml:"cluster.name"`
	Node_name        string `yaml:"node.name"`
	Discovery_type   string `yaml:"discovery.type"`
	Xpack_license    string `yaml:"xpack.license.self_generated.type"`
	Xpack_security   string `yaml:"xpack.security.enabled"`
	Xpack_enrollment string `yaml:"xpack.security.enrollment.enabled"`
	Xpack_http_ssl   string `yaml:"xpack.security.http.ssl.enabled"`
}

func CreateElasticsearch() Elasticsearch {
	config := Elasticsearch{
		Image:         CustomConfig.GetImage(Define.DOCKER_SERVICE_VDB_ELASTICSEARCH),
		ContainerName: "elasticsearch",
		Profiles:      []string{"elasticsearch"},
		Restart:       "always",
		Volumes: []string{
			"dify_es01_data:/usr/share/elasticsearch/data",
		},
		Environment: ElasticsearchEnvironment{
			ELASTIC_PASSWORD: "${ELASTICSEARCH_PASSWORD:-elastic}",
			Cluster_name:     "dify-es-cluster",
			Node_name:        "dify-es0",
			Discovery_type:   "single-node",
			Xpack_license:    "trial",
			Xpack_security:   "true",
			Xpack_enrollment: "false",
			Xpack_http_ssl:   "false",
		},
		Ports: []string{
			"${ELASTICSEARCH_PORT:-9200}:9200",
		},
	}

	healthCheckCmd, err := Fn.ConvertArrToCommand([]string{"CMD", "curl", "-s", "http://localhost:9200/_cluster/health?pretty"})
	if err == nil {
		config.Healthcheck = Define.HealthCheck{
			Test: healthCheckCmd,
		}
		config.Healthcheck.Interval = "30s"
		config.Healthcheck.Timeout = "10s"
		config.Healthcheck.Retries = 50
	}

	return config
}

type Kibana struct {
	Image         string             `yaml:"image"`
	ContainerName string             `yaml:"container_name"`
	Profiles      []string           `yaml:"profiles"`
	DependsOn     []string           `yaml:"depends_on"`
	Restart       string             `yaml:"restart"`
	Environment   KibanaEnvironment  `yaml:"environment"`
	Ports         []string           `yaml:"ports"`
	Healthcheck   Define.HealthCheck `yaml:"healthcheck"`
}

type KibanaEnvironment struct {
	XPACK_ENCRYPTEDSAVEDOBJECTS_ENCRYPTIONKEY string `yaml:"XPACK_ENCRYPTEDSAVEDOBJECTS_ENCRYPTIONKEY"`
	NO_PROXY                                  string `yaml:"NO_PROXY"`
	XPACK_SECURITY_ENABLED                    string `yaml:"XPACK_SECURITY_ENABLED"`
	XPACK_SECURITY_ENROLLMENT_ENABLED         string `yaml:"XPACK_SECURITY_ENROLLMENT_ENABLED"`
	XPACK_SECURITY_HTTP_SSL_ENABLED           string `yaml:"XPACK_SECURITY_HTTP_SSL_ENABLED"`
	XPACK_FLEET_ISAIRGAPPED                   string `yaml:"XPACK_FLEET_ISAIRGAPPED"`
	I18N_LOCALE                               string `yaml:"I18N_LOCALE"`
	SERVER_PORT                               string `yaml:"SERVER_PORT"`
	ELASTICSEARCH_HOSTS                       string `yaml:"ELASTICSEARCH_HOSTS"`
}

func CreateKibana() Kibana {
	config := Kibana{
		Image:         CustomConfig.GetImage(Define.DOCKER_SERVICE_VDB_ELASTICSEARCH_KIBANA),
		ContainerName: "kibana",
		Profiles:      []string{"elasticsearch"},
		DependsOn:     []string{"elasticsearch"},
		Restart:       "always",
		Environment: KibanaEnvironment{
			XPACK_ENCRYPTEDSAVEDOBJECTS_ENCRYPTIONKEY: "d1a66dfd-c4d3-4a0a-8290-2abcb83ab3aa",
			NO_PROXY:                          "localhost,127.0.0.1,elasticsearch,kibana",
			XPACK_SECURITY_ENABLED:            "true",
			XPACK_SECURITY_ENROLLMENT_ENABLED: "false",
			XPACK_SECURITY_HTTP_SSL_ENABLED:   "false",
			XPACK_FLEET_ISAIRGAPPED:           "true",
			I18N_LOCALE:                       "zh-CN",
			SERVER_PORT:                       "5601",
			ELASTICSEARCH_HOSTS:               "http://elasticsearch:9200",
		},
		Ports: []string{
			"${KIBANA_PORT:-5601}:5601",
		},
	}

	healthCheckCmd, err := Fn.ConvertArrToCommand([]string{"CMD-SHELL", "curl -s http://localhost:5601 >/dev/null || exit 1"})
	if err == nil {
		config.Healthcheck = Define.HealthCheck{
			Test: healthCheckCmd,
		}
		config.Healthcheck.Interval = "30s"
		config.Healthcheck.Timeout = "10s"
		config.Healthcheck.Retries = 3
	}

	return config
}
