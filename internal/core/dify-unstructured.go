package DifyCore

type Unstructured struct {
	Image    string   `yaml:"image"`
	Profiles []string `yaml:"profiles"`
	Restart  string   `yaml:"restart"`
	Volumes  []string `yaml:"volumes"`
}