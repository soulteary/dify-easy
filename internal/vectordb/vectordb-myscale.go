package VectorDB

type Myscale struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Profiles      []string `yaml:"profiles"`
	Restart       string   `yaml:"restart"`
	Tty           bool     `yaml:"tty"`
	Volumes       []string `yaml:"volumes"`
	Ports         []string `yaml:"ports"`
}
