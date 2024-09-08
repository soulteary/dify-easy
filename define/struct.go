package Define

type HealthCheck struct {
	Test        string `yaml:"test"`
	Interval    string `yaml:"interval,omitempty"`
	StartPeriod string `yaml:"start_period,omitempty"`
	Timeout     string `yaml:"timeout,omitempty"`
	Retries     int    `yaml:"retries,omitempty"`
}
