package VectorDB

import (
	Define "github.com/soulteary/dify-easy/define"
	CustomConfig "github.com/soulteary/dify-easy/internal/custom-config"
	Fn "github.com/soulteary/dify-easy/internal/fn"
)

type MilvusEtcd struct {
	ContainerName string                `yaml:"container_name"`
	Image         string                `yaml:"image"`
	Profiles      []string              `yaml:"profiles"`
	Environment   MilvusEtcdEnvironment `yaml:"environment"`
	Volumes       []string              `yaml:"volumes"`
	Command       string                `yaml:"command"`
	Healthcheck   Define.HealthCheck    `yaml:"healthcheck"`
	Networks      []string              `yaml:"networks"`
}

type MilvusEtcdEnvironment struct {
	ETCD_AUTO_COMPACTION_MODE      string `yaml:"ETCD_AUTO_COMPACTION_MODE"`
	ETCD_AUTO_COMPACTION_RETENTION string `yaml:"ETCD_AUTO_COMPACTION_RETENTION"`
	ETCD_QUOTA_BACKEND_BYTES       string `yaml:"ETCD_QUOTA_BACKEND_BYTES"`
	ETCD_SNAPSHOT_COUNT            string `yaml:"ETCD_SNAPSHOT_COUNT"`
}

func CreateMilvusEtcd() MilvusEtcd {
	config := MilvusEtcd{
		ContainerName: "milvus-etcd",
		Image:         CustomConfig.GetImage(Define.DOCKER_SERVICE_VDB_MILVUS_ETCD),
		Profiles:      []string{"milvus"},
		Environment: MilvusEtcdEnvironment{
			ETCD_AUTO_COMPACTION_MODE:      "${ETCD_AUTO_COMPACTION_MODE:-revision}",
			ETCD_AUTO_COMPACTION_RETENTION: "${ETCD_AUTO_COMPACTION_RETENTION:-1000}",
			ETCD_QUOTA_BACKEND_BYTES:       "${ETCD_QUOTA_BACKEND_BYTES:-4294967296}",
			ETCD_SNAPSHOT_COUNT:            "${ETCD_SNAPSHOT_COUNT:-50000}",
		},
		Volumes: []string{
			"./volumes/milvus/etcd:/etcd",
		},
		Networks: []string{"milvus"},
		Command:  `etcd -advertise-client-urls=http://127.0.0.1:2379 -listen-client-urls http://0.0.0.0:2379 --data-dir /etcd`,
	}

	healthCheckCmd, err := Fn.ConvertArrToCommand([]string{"CMD", "etcdctl", "endpoint", "health"})
	if err == nil {
		config.Healthcheck = Define.HealthCheck{
			Test: healthCheckCmd,
		}
		config.Healthcheck.Interval = "30s"
		config.Healthcheck.Timeout = "20s"
		config.Healthcheck.Retries = 3
	}

	return config
}

type MilvusMinio struct {
	ContainerName string                 `yaml:"container_name"`
	Image         string                 `yaml:"image"`
	Profiles      []string               `yaml:"profiles"`
	Environment   MilvusMinioEnvironment `yaml:"environment"`
	Volumes       []string               `yaml:"volumes"`
	Command       string                 `yaml:"command"`
	Healthcheck   Define.HealthCheck     `yaml:"healthcheck"`
	Networks      []string               `yaml:"networks"`
}

type MilvusMinioEnvironment struct {
	MINIO_ACCESS_KEY string `yaml:"MINIO_ACCESS_KEY"`
	MINIO_SECRET_KEY string `yaml:"MINIO_SECRET_KEY"`
}

func CreateMilvusMinio() MilvusMinio {
	config := MilvusMinio{
		ContainerName: "milvus-minio",
		Image:         CustomConfig.GetImage(Define.DOCKER_SERVICE_VDB_MILVUS_MINIO),
		Profiles:      []string{"milvus"},
		Environment: MilvusMinioEnvironment{
			MINIO_ACCESS_KEY: "${MINIO_ACCESS_KEY:-minioadmin}",
			MINIO_SECRET_KEY: "${MINIO_SECRET_KEY:-minioadmin}",
		},
		Volumes: []string{
			"./volumes/milvus/minio:/minio_data",
		},
		Networks: []string{"milvus"},
		Command:  `minio server /minio_data --console-address ":9001"`,
	}

	healthCheckCmd, err := Fn.ConvertArrToCommand([]string{"CMD", "curl", "-f", "http://localhost:9000/minio/health/live"})
	if err == nil {
		config.Healthcheck = Define.HealthCheck{
			Test: healthCheckCmd,
		}
		config.Healthcheck.Interval = "30s"
		config.Healthcheck.Timeout = "20s"
		config.Healthcheck.Retries = 3
	}
	return config
}

type MilvusStandalone struct {
	ContainerName string             `yaml:"container_name"`
	Image         string             `yaml:"image"`
	Profiles      []string           `yaml:"profiles"`
	Command       string             `yaml:"command"`
	Environment   MilvusEnvironment  `yaml:"environment"`
	Volumes       []string           `yaml:"volumes"`
	Healthcheck   Define.HealthCheck `yaml:"healthcheck"`
	DependsOn     []string           `yaml:"depends_on"`
	Networks      []string           `yaml:"networks"`
	Ports         []string           `yaml:"ports"`
}

type MilvusEnvironment struct {
	ETCD_ENDPOINTS                        string `yaml:"ETCD_ENDPOINTS"`
	MINIO_ADDRESS                         string `yaml:"MINIO_ADDRESS"`
	Common_Security_Authorization_Enabled string `yaml:"common.security.authorizationEnabled"`
}

func CreateMilvusStandalone() MilvusStandalone {
	config := MilvusStandalone{
		ContainerName: "milvus-standalone",
		Image:         CustomConfig.GetImage(Define.DOCKER_SERVICE_VDB_MILVUS),
		Profiles:      []string{"milvus"},
		Environment: MilvusEnvironment{
			ETCD_ENDPOINTS:                        "${ETCD_ENDPOINTS:-etcd:2379}",
			MINIO_ADDRESS:                         "${MINIO_ADDRESS:-minio:9000}",
			Common_Security_Authorization_Enabled: "${MILVUS_AUTHORIZATION_ENABLED:-true}",
		},
		Volumes: []string{
			"./volumes/milvus/milvus:/var/lib/milvus",
		},
		DependsOn: []string{"etcd", "minio"},
		Networks:  []string{"milvus"},
		Ports:     []string{"19530:19530", "9091:9091"},
	}

	healthCheckCmd, err := Fn.ConvertArrToCommand([]string{"CMD", "curl", "-f", "http://localhost:9091/healthz"})
	if err == nil {
		config.Healthcheck = Define.HealthCheck{
			Test: healthCheckCmd,
		}
		config.Healthcheck.Interval = "30s"
		config.Healthcheck.StartPeriod = "90s"
		config.Healthcheck.Timeout = "20s"
		config.Healthcheck.Retries = 3
	}

	command, err := Fn.ConvertArrToCommand([]string{"milvus", "run", "standalone"})
	if err == nil {
		config.Command = command
	}

	return config
}
