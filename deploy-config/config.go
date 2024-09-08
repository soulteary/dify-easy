package DeployConfig

import (
	"bytes"
	"fmt"

	CustomConfig "github.com/soulteary/dify-easy/custom-config"
	DifyCore "github.com/soulteary/dify-easy/internal/core"
	DifyNetwork "github.com/soulteary/dify-easy/internal/network"
	VectorDB "github.com/soulteary/dify-easy/internal/vectordb"
	DifyVolume "github.com/soulteary/dify-easy/internal/volume"
	"gopkg.in/yaml.v3"
)

func Create() DifyDeploy {
	config := DifyDeploy{}

	config.XSharedEnv = CustomConfig.GetXSharedEnv()

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

func (config DifyDeploy) ToString() string {
	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	err := encoder.Encode(config)
	if err != nil {
		fmt.Printf("Error encoding YAML: %v\n", err)
		return ""
	}
	return fixYamlFormat(buf.String())
}
