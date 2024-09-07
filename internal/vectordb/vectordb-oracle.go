package VectorDB

type Oracle struct {
	Image       string        `yaml:"image"`
	Profiles    []string      `yaml:"profiles"`
	Restart     string        `yaml:"restart"`
	Volumes     []interface{} `yaml:"volumes"`
	Environment []string      `yaml:"environment"`
}
