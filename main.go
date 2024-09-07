package main

import (
	"fmt"
	"log"
	"os"

	"github.com/soulteary/dify-easy/deploy"
	"github.com/soulteary/dify-easy/fn"

	"gopkg.in/yaml.v3"
)

func main() {
	config := Deploy.CreateConfig()
	yamlData, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// Fix the YAML
	output := Fn.FixYAML(yamlData)

	// Print the YAML
	fmt.Printf("--- YAML ---\n%s\n", output)

	os.WriteFile("docker-compose.test.yaml", []byte(output), 0644)

}
