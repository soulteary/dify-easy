package DeployConfig

import (
	Define "github.com/soulteary/dify-easy/define"
	DifyCore "github.com/soulteary/dify-easy/internal/core"
	DifyNetwork "github.com/soulteary/dify-easy/internal/network"
	VectorDB "github.com/soulteary/dify-easy/internal/vectordb"
	DifyVolume "github.com/soulteary/dify-easy/internal/volume"
)

type DifyDeploy struct {
	Define.XSharedEnv `yaml:"x-shared-env,omitempty"`

	Services struct {
		DifyCore.API       `yaml:"api"`
		DifyCore.Worker    `yaml:"worker"`
		DifyCore.Web       `yaml:"web"`
		DifyCore.DB        `yaml:"db"`
		DifyCore.Redis     `yaml:"redis"`
		DifyCore.Sandbox   `yaml:"sandbox"`
		DifyCore.SsrfProxy `yaml:"ssrf_proxy"`
		DifyCore.Certbot   `yaml:"certbot"`
		DifyCore.Nginx     `yaml:"nginx"`

		VectorDB.Weaviate `yaml:"weaviate,omitempty"`

		VectorDB.Qdrant `yaml:"qdrant,omitempty"`

		VectorDB.Pgvector `yaml:"pgvector,omitempty"`

		VectorDB.PgvectoRs `yaml:"pgvecto-rs,omitempty"`

		VectorDB.Chroma `yaml:"chroma,omitempty"`

		VectorDB.Oracle `yaml:"oracle,omitempty"`

		VectorDB.MilvusEtcd       `yaml:"etcd,omitempty"`
		VectorDB.MilvusMinio      `yaml:"minio,omitempty"`
		VectorDB.MilvusStandalone `yaml:"milvus-standalone,omitempty"`

		VectorDB.Opensearch           `yaml:"opensearch,omitempty"`
		VectorDB.OpensearchDashboards `yaml:"opensearch-dashboards,omitempty"`

		VectorDB.Myscale `yaml:"myscale,omitempty"`

		VectorDB.Elasticsearch `yaml:"elasticsearch,omitempty"`
		VectorDB.Kibana        `yaml:"kibana,omitempty"`

		DifyCore.Unstructured `yaml:"unstructured,omitempty"`
	} `yaml:"services"`

	DifyNetwork.Networks `yaml:"networks"`

	DifyVolume.Volumes `yaml:"volumes"`
}
