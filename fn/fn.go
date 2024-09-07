package Fn

import "strings"

func FixYAML(raw string) string {
	output := raw

	output = strings.ReplaceAll(output, `x-shared-env:`, "x-shared-env: &shared-api-worker-env")
	output = strings.ReplaceAll(output, "<<: '*shared-api-worker-env'", "<<: *shared-api-worker-env")

	output = strings.ReplaceAll(output, "command: |-\n      >", "command: >")

	output = strings.ReplaceAll(output, `: keep-empty`, ":")
	return output
}
