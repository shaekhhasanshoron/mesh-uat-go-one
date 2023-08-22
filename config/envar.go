package config

import (
	"log"
	"os"
	"strings"
)

var MeshUatGoTwoEndpoint string
var RunMode string
var ServerPort string

func InitEnvironmentVariables() {
	RunMode = strings.TrimSpace(os.Getenv("RUN_MODE"))
	if RunMode == "" {
		RunMode = "DEVELOP"
	}

	ServerPort = strings.TrimSpace(os.Getenv("SERVER_PORT"))
	if ServerPort == "" {
		ServerPort = "8080"
	}

	MeshUatGoTwoEndpoint = strings.TrimSpace(os.Getenv("MESH_UAT_GO_TWO_ENDPOINT"))
	if MeshUatGoTwoEndpoint == "" {
		MeshUatGoTwoEndpoint = "http://localhost:8081"
	}

	log.Println("Run Mode: " + RunMode)
	log.Println("Server Port: " + ServerPort)
	log.Println("MeshUatGoTwo Endpoint: " + MeshUatGoTwoEndpoint)
}
