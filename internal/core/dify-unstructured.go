package DifyCore

type Unstructured struct {
	Image    string   `yaml:"image"`
	Profiles []string `yaml:"profiles"`
	Restart  string   `yaml:"restart"`
	Volumes  []string `yaml:"volumes"`
}

func CreateUnstructured() Unstructured {
	return Unstructured{
		Image:    "downloads.unstructured.io/unstructured-io/unstructured-api:latest",
		Profiles: []string{"unstructured"},
		Restart:  "always",
		Volumes: []string{
			"./volumes/unstructured:/app/data",
		},
	}
}
