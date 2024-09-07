package VectorDB

type Opensearch struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Profiles      []string `yaml:"profiles"`
	Environment   []string `yaml:"environment"`
	Ulimits       struct {
		Memlock struct {
			Soft string `yaml:"soft"`
			Hard string `yaml:"hard"`
		} `yaml:"memlock"`
		Nofile struct {
			Soft string `yaml:"soft"`
			Hard string `yaml:"hard"`
		} `yaml:"nofile"`
	} `yaml:"ulimits"`
	Volumes  []string `yaml:"volumes"`
	Networks []string `yaml:"networks"`
}
type OpensearchDashboards struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Profiles      []string `yaml:"profiles"`
	Environment   struct {
		OPENSEARCHHOSTS string `yaml:"OPENSEARCH_HOSTS"`
	} `yaml:"environment"`
	Volumes   []string `yaml:"volumes"`
	Networks  []string `yaml:"networks"`
	DependsOn []string `yaml:"depends_on"`
}
