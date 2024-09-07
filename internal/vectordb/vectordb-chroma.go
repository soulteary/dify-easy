package VectorDB

type Chroma struct {
	Image       string   `yaml:"image"`
	Profiles    []string `yaml:"profiles"`
	Restart     string   `yaml:"restart"`
	Volumes     []string `yaml:"volumes"`
	Environment struct {
		CHROMASERVERAUTHNCREDENTIALS string `yaml:"CHROMA_SERVER_AUTHN_CREDENTIALS"`
		CHROMASERVERAUTHNPROVIDER    string `yaml:"CHROMA_SERVER_AUTHN_PROVIDER"`
		ISPERSISTENT                 string `yaml:"IS_PERSISTENT"`
	} `yaml:"environment"`
}
