package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/soulteary/dify-easy/deploy"
	"github.com/soulteary/dify-easy/fn"

	"gopkg.in/yaml.v3"
)

func main() {
	config := Deploy.CreateConfig()

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	err := encoder.Encode(config)
	if err != nil {
		fmt.Printf("Error encoding YAML: %v\n", err)
		return
	}
	yamlData := buf.String()

	// Fix the YAML
	output := Fn.FixYAML(yamlData)

	// Print the YAML
	fmt.Printf("--- YAML ---\n%s\n", output)

	os.WriteFile("./docker-compose.yml", []byte(output), 0644)
}
