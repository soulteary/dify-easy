package VectorDB

type PgvectoRs struct {
	Image       string   `yaml:"image"`
	Profiles    []string `yaml:"profiles"`
	Restart     string   `yaml:"restart"`
	Environment struct {
		PGUSER           string `yaml:"PGUSER"`
		POSTGRESPASSWORD string `yaml:"POSTGRES_PASSWORD"`
		POSTGRESDB       string `yaml:"POSTGRES_DB"`
		PGDATA           string `yaml:"PGDATA"`
	} `yaml:"environment"`
	Volumes     []string `yaml:"volumes"`
	Healthcheck struct {
		Test     []string `yaml:"test"`
		Interval string   `yaml:"interval"`
		Timeout  string   `yaml:"timeout"`
		Retries  int      `yaml:"retries"`
	} `yaml:"healthcheck"`
}
