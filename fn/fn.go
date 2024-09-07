package Fn

import (
	"encoding/json"
	"regexp"
	"strings"
)

func FixYAML(raw string) string {
	output := raw

	output = strings.ReplaceAll(output, `x-shared-env:`, "x-shared-env: &shared-api-worker-env")
	output = strings.ReplaceAll(output, "<<: '*shared-api-worker-env'", "<<: *shared-api-worker-env")

	// fix muti-line command
	output = strings.ReplaceAll(output, "command: |-\n      >", "command: >")
	// fix healthcheck command
	var re = regexp.MustCompile(`(?m)(test:\s)\'(\[.*?\])'`)
	output = re.ReplaceAllString(output, "$1$2")

	output = strings.ReplaceAll(output, `: keep-empty`, ":")
	return output
}

func GetHealthCheckCMD(arr []string) (string, error) {
	jsonData, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}
