package VectorDB

type Etcd struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Profiles      []string `yaml:"profiles"`
	Environment   []string `yaml:"environment"`
	Volumes       []string `yaml:"volumes"`
	Command       string   `yaml:"command"`
	Healthcheck   struct {
		Test     []string `yaml:"test"`
		Interval string   `yaml:"interval"`
		Timeout  string   `yaml:"timeout"`
		Retries  int      `yaml:"retries"`
	} `yaml:"healthcheck"`
	Networks []string `yaml:"networks"`
}
type Minio struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Profiles      []string `yaml:"profiles"`
	Environment   struct {
		MINIOACCESSKEY string `yaml:"MINIO_ACCESS_KEY"`
		MINIOSECRETKEY string `yaml:"MINIO_SECRET_KEY"`
	} `yaml:"environment"`
	Volumes     []string `yaml:"volumes"`
	Command     string   `yaml:"command"`
	Healthcheck struct {
		Test     []string `yaml:"test"`
		Interval string   `yaml:"interval"`
		Timeout  string   `yaml:"timeout"`
		Retries  int      `yaml:"retries"`
	} `yaml:"healthcheck"`
	Networks []string `yaml:"networks"`
}
type MilvusStandalone struct {
	ContainerName string   `yaml:"container_name"`
	Image         string   `yaml:"image"`
	Profiles      []string `yaml:"profiles"`
	Command       []string `yaml:"command"`
	Environment   struct {
		ETCDENDPOINTS                      string `yaml:"ETCD_ENDPOINTS"`
		MINIOADDRESS                       string `yaml:"MINIO_ADDRESS"`
		CommonSecurityAuthorizationEnabled string `yaml:"common.security.authorizationEnabled"`
	} `yaml:"environment"`
	Volumes     []string `yaml:"volumes"`
	Healthcheck struct {
		Test        []string `yaml:"test"`
		Interval    string   `yaml:"interval"`
		StartPeriod string   `yaml:"start_period"`
		Timeout     string   `yaml:"timeout"`
		Retries     int      `yaml:"retries"`
	} `yaml:"healthcheck"`
	DependsOn []string `yaml:"depends_on"`
	Networks  []string `yaml:"networks"`
}
