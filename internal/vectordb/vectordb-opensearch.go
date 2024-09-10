package VectorDB

import (
	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	Define "github.com/soulteary/dify-easy/define"
)

type Opensearch struct {
	ContainerName string                `yaml:"container_name"`
	Image         string                `yaml:"image"`
	Profiles      []string              `yaml:"profiles"`
	Environment   OpensearchEnvironment `yaml:"environment"`
	Ulimits       Ulimits               `yaml:"ulimits"`
	Volumes       []string              `yaml:"volumes"`
	Networks      []string              `yaml:"networks"`
}

type Memlock struct {
	Soft string `yaml:"soft"`
	Hard string `yaml:"hard"`
}

type Nofile struct {
	Soft string `yaml:"soft"`
	Hard string `yaml:"hard"`
}

type Ulimits struct {
	Memlock Memlock `yaml:"memlock"`
	Nofile  Nofile  `yaml:"nofile"`
}

type OpensearchEnvironment struct {
	Discovery_type                    string `yaml:"discovery.type"`
	Bootstrap_memory_lock             string `yaml:"bootstrap.memory_lock"`
	OPENSEARCH_JAVA_OPTS              string `yaml:"OPENSEARCH_JAVA_OPTS"`
	OPENSEARCH_INITIAL_ADMIN_PASSWORD string `yaml:"OPENSEARCH_INITIAL_ADMIN_PASSWORD"`
}

func CreateOpensearch() Opensearch {
	config := Opensearch{
		ContainerName: "opensearch",
		Image:         "opensearchproject/opensearch:latest",
		Profiles:      []string{"opensearch"},
		Environment: OpensearchEnvironment{
			Discovery_type:                    "${OPENSEARCH_DISCOVERY_TYPE:-single-node}",
			Bootstrap_memory_lock:             "${OPENSEARCH_BOOTSTRAP_MEMORY_LOCK:-true}",
			OPENSEARCH_JAVA_OPTS:              "-Xms${OPENSEARCH_JAVA_OPTS_MIN:-512m} -Xmx${OPENSEARCH_JAVA_OPTS_MAX:-1024m}",
			OPENSEARCH_INITIAL_ADMIN_PASSWORD: "${OPENSEARCH_INITIAL_ADMIN_PASSWORD:-Qazwsxedc!@#123}",
		},
		Ulimits: Ulimits{
			Memlock: Memlock{
				Soft: "${OPENSEARCH_MEMLOCK_SOFT:--1}",
				Hard: "${OPENSEARCH_MEMLOCK_HARD:--1}",
			},
			Nofile: Nofile{
				Soft: "${OPENSEARCH_NOFILE_SOFT:-65536}",
				Hard: "${OPENSEARCH_NOFILE_HARD:-65536}",
			},
		},
		Volumes: []string{
			"./volumes/opensearch/data:/usr/share/opensearch/data",
		},
		Networks: []string{"opensearch-net"},
	}
	return config
}

type OpensearchDashboards struct {
	ContainerName string                          `yaml:"container_name"`
	Image         string                          `yaml:"image"`
	Profiles      []string                        `yaml:"profiles"`
	Environment   OpensearchDashboardsEnvironment `yaml:"environment"`
	Volumes       []string                        `yaml:"volumes"`
	Networks      []string                        `yaml:"networks"`
	DependsOn     []string                        `yaml:"depends_on"`
}

type OpensearchDashboardsEnvironment struct {
	OPENSEARCH_HOSTS string `yaml:"OPENSEARCH_HOSTS"`
}

func CreateOpenSearchDashboards() OpensearchDashboards {
	config := OpensearchDashboards{
		ContainerName: "opensearch-dashboards",
		Image:         CustomConfig.GetImage(Define.DOCKER_SERVICE_VDB_OPENSEARCH),
		Profiles:      []string{"opensearch"},
		Environment: OpensearchDashboardsEnvironment{
			OPENSEARCH_HOSTS: `["https://opensearch:9200"]`,
		},
		Volumes: []string{
			"./volumes/opensearch/opensearch_dashboards.yml:/usr/share/opensearch-dashboards/config/opensearch_dashboards.yml",
		},
		Networks:  []string{"opensearch-net"},
		DependsOn: []string{"opensearch"},
	}
	return config
}
