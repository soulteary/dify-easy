package Fn

import (
	"bytes"
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
	// fix entrypoint command
	output = regexp.MustCompile(`(?m)(entrypoint:\s)\|\n\s+(\[.*\])`).ReplaceAllString(output, "$1$2")
	// fix command
	output = regexp.MustCompile(`(?m)(command:\s)\|\n\s+(\[.*\])`).ReplaceAllString(output, "$1$2")
	// fix healthcheck command
	output = regexp.MustCompile(`(?m)(test:\s)\|\n\s+(\[.*\])`).ReplaceAllString(output, "$1$2")

	for _, match := range regexp.MustCompile(`(?m)^\s+test:\s\[(.+)\]$`).FindAllStringSubmatch(output, -1) {
		sub := regexp.MustCompile(`(?m),`).ReplaceAllString(match[0], ", ")
		sub = regexp.MustCompile(`(?m)\[(.*)\]`).ReplaceAllString(sub, "[ $1 ]")
		output = strings.ReplaceAll(output, match[0], sub)
	}

	// fix ports
	output = regexp.MustCompile(`(?m)(-\s)'"(.*)"'$`).ReplaceAllString(output, "$1\"$2\"")

	output = strings.ReplaceAll(output, `: keep-empty`, ":")
	return output
}

func ConvertArrToCommand(arr []string) (string, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(arr)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
