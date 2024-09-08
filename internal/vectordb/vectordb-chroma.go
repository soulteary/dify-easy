package VectorDB

type Chroma struct {
	Image       string            `yaml:"image"`
	Profiles    []string          `yaml:"profiles"`
	Restart     string            `yaml:"restart"`
	Volumes     []string          `yaml:"volumes"`
	Environment ChromaEnvironment `yaml:"environment"`
}

type ChromaEnvironment struct {
	CHROMA_SERVER_AUTHN_CREDENTIALS string `yaml:"CHROMA_SERVER_AUTHN_CREDENTIALS"`
	CHROMA_SERVER_AUTHN_PROVIDER    string `yaml:"CHROMA_SERVER_AUTHN_PROVIDER"`
	IS_PERSISTENT                   string `yaml:"IS_PERSISTENT"`
}

func CreateChroma() Chroma {
	config := Chroma{
		Image:    "ghcr.io/chroma-core/chroma:0.5.1",
		Profiles: []string{"chroma"},
		Restart:  "always",
		Volumes: []string{
			"./volumes/chroma:/chroma/chroma",
		},
		Environment: ChromaEnvironment{
			CHROMA_SERVER_AUTHN_CREDENTIALS: "${CHROMA_SERVER_AUTHN_CREDENTIALS:-difyai123456}",
			CHROMA_SERVER_AUTHN_PROVIDER:    "${CHROMA_SERVER_AUTHN_PROVIDER:-chromadb.auth.token_authn.TokenAuthenticationServerProvider}",
			IS_PERSISTENT:                   "${CHROMA_IS_PERSISTENT:-TRUE}",
		},
	}
	return config
}
