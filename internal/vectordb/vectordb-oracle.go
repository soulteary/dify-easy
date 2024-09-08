package VectorDB

type Oracle struct {
	Image       string            `yaml:"image"`
	Profiles    []string          `yaml:"profiles"`
	Restart     string            `yaml:"restart"`
	Volumes     []interface{}     `yaml:"volumes"`
	Environment OracleEnvironment `yaml:"environment"`
}

type OracleEnvironment struct {
	ORACLE_PWD          string `yaml:"ORACLE_PWD"`
	ORACLE_CHARACTERSET string `yaml:"ORACLE_CHARACTERSET"`
}

func CreateOracle() Oracle {
	config := Oracle{
		Image:    "container-registry.oracle.com/database/free:latest",
		Profiles: []string{"oracle"},
		Restart:  "always",
		Volumes: []interface{}{
			map[string]interface{}{
				"type":   "volume",
				"source": "oradata",
				"target": "/opt/oracle/oradata",
			},
			"./startupscripts:/opt/oracle/scripts/startup",
		},
		Environment: OracleEnvironment{
			ORACLE_PWD:          "${ORACLE_PWD:-Dify123456}",
			ORACLE_CHARACTERSET: "${ORACLE_CHARACTERSET:-AL32UTF8}",
		},
	}
	return config
}
