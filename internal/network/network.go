package DifyNetwork

type Networks struct {
	SsrfProxyNetwork Network `yaml:"ssrf_proxy_network"`
	Milvus           Network `yaml:"milvus"`
	OpensearchNet    Network `yaml:"opensearch-net"`
}

type Network struct {
	Driver   string `yaml:"driver"`
	Internal bool   `yaml:"internal,omitempty"`
}

func CreateNetworks() Networks {
	networks := Networks{}
	networks.SsrfProxyNetwork = Network{
		Driver:   "bridge",
		Internal: true,
	}

	networks.Milvus = Network{
		Driver: "bridge",
	}

	networks.OpensearchNet = Network{
		Driver:   "bridge",
		Internal: true,
	}

	return networks
}
