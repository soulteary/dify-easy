package DifyVolume

type Volumes struct {
	Oradata      interface{} `yaml:"oradata"`
	DifyEs01Data interface{} `yaml:"dify_es01_data"`
}

func CreateVolumes() Volumes {
	return Volumes{
		Oradata:      "keep-empty",
		DifyEs01Data: "keep-empty",
	}
}
