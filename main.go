package main

import (
	"fmt"
	"os"

	"github.com/soulteary/dify-easy/deploy-config"
)

func main() {
	config := DeployConfig.Create()

	output := config.ToString()

	fmt.Printf("--- YAML ---\n%s\n", output)

	os.WriteFile("../test-dirs/docker-compose.yml", []byte(output), 0644)
}
