package VectorDB

type Weaviate struct {
	Image       string   `yaml:"image"`
	Profiles    []string `yaml:"profiles"`
	Restart     string   `yaml:"restart"`
	Volumes     []string `yaml:"volumes"`
	Environment struct {
		PERSISTENCEDATAPATH                  string `yaml:"PERSISTENCE_DATA_PATH"`
		QUERYDEFAULTSLIMIT                   string `yaml:"QUERY_DEFAULTS_LIMIT"`
		AUTHENTICATIONANONYMOUSACCESSENABLED string `yaml:"AUTHENTICATION_ANONYMOUS_ACCESS_ENABLED"`
		DEFAULTVECTORIZERMODULE              string `yaml:"DEFAULT_VECTORIZER_MODULE"`
		CLUSTERHOSTNAME                      string `yaml:"CLUSTER_HOSTNAME"`
		AUTHENTICATIONAPIKEYENABLED          string `yaml:"AUTHENTICATION_APIKEY_ENABLED"`
		AUTHENTICATIONAPIKEYALLOWEDKEYS      string `yaml:"AUTHENTICATION_APIKEY_ALLOWED_KEYS"`
		AUTHENTICATIONAPIKEYUSERS            string `yaml:"AUTHENTICATION_APIKEY_USERS"`
		AUTHORIZATIONADMINLISTENABLED        string `yaml:"AUTHORIZATION_ADMINLIST_ENABLED"`
		AUTHORIZATIONADMINLISTUSERS          string `yaml:"AUTHORIZATION_ADMINLIST_USERS"`
	} `yaml:"environment"`
}
