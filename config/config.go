package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct{
	Version string
	ServiceName string
	HttpPort int
}

var configuration Config

func loadConfig(){
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Failed to load the env variable: %v", err)
		return
	}
	version := os.Getenv("VERSION")
	if version == ""{
		log.Println("Version is required")
		return
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == ""{
		log.Println("Service Name is required")
		return
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == ""{
		log.Fatal("Http Port is required")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil{
		fmt.Println("port must be number")
		os.Exit(1)
	}

	configuration = Config{
		Version: version,
		ServiceName: serviceName,
		HttpPort: port,
	}
}

func GetConfig() Config{
	loadConfig()
	return configuration
}