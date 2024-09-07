package VectorDB

type Elasticsearch struct {
	Image         string   `yaml:"image"`
	ContainerName string   `yaml:"container_name"`
	Profiles      []string `yaml:"profiles"`
	Restart       string   `yaml:"restart"`
	Volumes       []string `yaml:"volumes"`
	Environment   []string `yaml:"environment"`
	Ports         []string `yaml:"ports"`
	Healthcheck   struct {
		Test     []string `yaml:"test"`
		Interval string   `yaml:"interval"`
		Timeout  string   `yaml:"timeout"`
		Retries  int      `yaml:"retries"`
	} `yaml:"healthcheck"`
}
type Kibana struct {
	Image         string   `yaml:"image"`
	ContainerName string   `yaml:"container_name"`
	Profiles      []string `yaml:"profiles"`
	DependsOn     []string `yaml:"depends_on"`
	Restart       string   `yaml:"restart"`
	Environment   []string `yaml:"environment"`
	Ports         []string `yaml:"ports"`
	Healthcheck   struct {
		Test     []string `yaml:"test"`
		Interval string   `yaml:"interval"`
		Timeout  string   `yaml:"timeout"`
		Retries  int      `yaml:"retries"`
	} `yaml:"healthcheck"`
}
