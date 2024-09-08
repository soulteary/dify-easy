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

	config.Services.MilvusEtcd = VectorDB.CreateMilvusEtcd()
	config.Services.MilvusMinio = VectorDB.CreateMilvusMinio()
	config.Services.MilvusStandalone = VectorDB.CreateMilvusStandalone()

	config.Services.Opensearch = VectorDB.CreateOpensearch()
	config.Services.OpensearchDashboards = VectorDB.CreateOpenSearchDashboards()

	config.Services.Myscale = VectorDB.CreateMyscale()

	config.Services.Elasticsearch = VectorDB.CreateElasticsearch()
	config.Services.Kibana = VectorDB.CreateKibana()

	config.Services.Unstructured = DifyCore.CreateUnstructured()

	config.Networks = DifyNetwork.CreateNetworks()
	config.Volumes = DifyVolume.CreateVolumes()

	return config
}
