package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/CatGitBon/auth-service/internal/config"
)

func main() {

	result, err := run()

	fmt.Println(result, err)
}

func run() (string, error) {

	configPath := flag.String("config", "./config", "path to the config file")

	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		return "", fmt.Errorf("error loading config: %v", err)
	}

	return "", errors.New("khoih")
}
