package VectorDB

type Qdrant struct {
	Image       string   `yaml:"image"`
	Profiles    []string `yaml:"profiles"`
	Restart     string   `yaml:"restart"`
	Volumes     []string `yaml:"volumes"`
	Environment struct {
		QDRANTAPIKEY string `yaml:"QDRANT_API_KEY"`
	} `yaml:"environment"`
}
