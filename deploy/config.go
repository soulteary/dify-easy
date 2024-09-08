package Deploy

import (
	DifyCore "github.com/soulteary/dify-easy/internal/core"
	DifyNetwork "github.com/soulteary/dify-easy/internal/network"
	VectorDB "github.com/soulteary/dify-easy/internal/vectordb"
	DifyVolume "github.com/soulteary/dify-easy/internal/volume"
)

type DifyDeploy struct {
	DifyCore.XSharedEnv `yaml:"x-shared-env,omitempty"`

	Services struct {
		DifyCore.API          `yaml:"api"`
		DifyCore.Worker       `yaml:"worker"`
		DifyCore.Web          `yaml:"web"`
		DifyCore.DB           `yaml:"db"`
		DifyCore.Redis        `yaml:"redis"`
		DifyCore.Sandbox      `yaml:"sandbox"`
		DifyCore.SsrfProxy    `yaml:"ssrf_proxy"`
		DifyCore.Certbot      `yaml:"certbot"`
		DifyCore.Nginx        `yaml:"nginx"`
		DifyCore.Unstructured `yaml:"unstructured,omitempty"`

		VectorDB.Weaviate `yaml:"weaviate,omitempty"`

		VectorDB.Qdrant `yaml:"qdrant,omitempty"`

		VectorDB.Pgvector `yaml:"pgvector,omitempty"`

		VectorDB.PgvectoRs `yaml:"pgvecto-rs,omitempty"`

		VectorDB.Chroma `yaml:"chroma,omitempty"`

		VectorDB.Oracle `yaml:"oracle,omitempty"`

		VectorDB.Opensearch           `yaml:"opensearch,omitempty"`
		VectorDB.OpensearchDashboards `yaml:"opensearch-dashboards,omitempty"`

		VectorDB.Myscale `yaml:"myscale,omitempty"`

		VectorDB.Elasticsearch `yaml:"elasticsearch,omitempty"`
		VectorDB.Kibana        `yaml:"kibana,omitempty"`

		VectorDB.Etcd             `yaml:"etcd,omitempty"`
		VectorDB.Minio            `yaml:"minio,omitempty"`
		VectorDB.MilvusStandalone `yaml:"milvus-standalone,omitempty"`
	} `yaml:"services"`

	DifyNetwork.Networks `yaml:"networks"`

	DifyVolume.Volumes `yaml:"volumes"`
}

func CreateConfig() DifyDeploy {
	config := DifyDeploy{}

	config.XSharedEnv = DifyCore.CreateXSharedEnv()

	config.Services.API = DifyCore.CreateDifyAPI()
	config.Services.Worker = DifyCore.CreateDifyWorker()
	config.Services.Web = DifyCore.CreateDifyWeb()
	config.Services.DB = DifyCore.CreateDifyDB()
	config.Services.Redis = DifyCore.CreateDifyRedis()
	config.Services.Sandbox = DifyCore.CreateDifySandbox()
	config.Services.Certbot = DifyCore.CreateDifyCertbot()
	config.Services.Nginx = DifyCore.CreateDifyNginx()
	config.Services.SsrfProxy = DifyCore.CreateDifySsrfProxy()

	config.Services.Weaviate = VectorDB.CreateWeaviate()
	config.Services.Qdrant = VectorDB.CreateQdrant()
	config.Services.Pgvector = VectorDB.CreatePgvector()
	config.Services.PgvectoRs = VectorDB.CreatePgvectorRs()
	config.Services.Chroma = VectorDB.CreateChroma()
	config.Services.Oracle = VectorDB.CreateOracle()

	config.Networks = DifyNetwork.CreateNetworks()
	config.Volumes = DifyVolume.CreateVolumes()

	return config
}
