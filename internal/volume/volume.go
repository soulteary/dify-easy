package DifyVolume

type Volumes struct {
	Oradata      interface{} `yaml:"oradata"`
	DifyEs01Data interface{} `yaml:"dify_es01_data"`
}

func CreateVolumes() Volumes {
	volumes := Volumes{}
	volumes.Oradata = "keep-empty"
	volumes.DifyEs01Data = "keep-empty"
	return volumes
}
