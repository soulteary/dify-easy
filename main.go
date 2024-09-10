package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	DeployConfig "github.com/soulteary/dify-easy/deploy-config"
	"github.com/soulteary/dify-easy/internal/server"
	"github.com/soulteary/dify-easy/internal/version"
)

//go:embed web
var EmbedFS embed.FS

func main() {
	log.Println("Dify Easy", version.Version)
	server.Launch(EmbedFS)

	config := DeployConfig.Create()

	output := config.ToString()

	fmt.Printf("--- YAML ---\n%s\n", output)

	os.WriteFile("docker-compose.yml", []byte(output), 0644)
}
