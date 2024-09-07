package DifyNetwork

type Networks struct {
	SsrfProxyNetwork struct {
		Driver   string `yaml:"driver"`
		Internal bool   `yaml:"internal"`
	} `yaml:"ssrf_proxy_network"`
	Milvus struct {
		Driver string `yaml:"driver"`
	} `yaml:"milvus"`
	OpensearchNet struct {
		Driver   string `yaml:"driver"`
		Internal bool   `yaml:"internal"`
	} `yaml:"opensearch-net"`
}

func CreateNetworks() Networks {
	return Networks{
		SsrfProxyNetwork: struct {
			Driver   string `yaml:"driver"`
			Internal bool   `yaml:"internal"`
		}{
			Driver:   "bridge",
			Internal: true,
		},
		Milvus: struct {
			Driver string `yaml:"driver"`
		}{
			Driver: "bridge",
		},
		OpensearchNet: struct {
			Driver   string `yaml:"driver"`
			Internal bool   `yaml:"internal"`
		}{
			Driver:   "bridge",
			Internal: true,
		},
	}
}
