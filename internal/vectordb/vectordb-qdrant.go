package VectorDB

type Qdrant struct {
	Image       string            `yaml:"image"`
	Profiles    []string          `yaml:"profiles"`
	Restart     string            `yaml:"restart"`
	Volumes     []string          `yaml:"volumes"`
	Environment QdrantEnvironment `yaml:"environment"`
}

type QdrantEnvironment struct {
	QDRANT_API_KEY string `yaml:"QDRANT_API_KEY"`
}

func CreateQdrant() Qdrant {
	config := Qdrant{
		Image: "langgenius/qdrant:v1.7.3",
		Profiles: []string{
			"qdrant",
		},
		Restart: "always",
		Volumes: []string{
			"./volumes/qdrant:/qdrant/storage",
		},
		Environment: QdrantEnvironment{
			QDRANT_API_KEY: "${QDRANT_API_KEY:-difyai123456}",
		},
	}

	return config
}
