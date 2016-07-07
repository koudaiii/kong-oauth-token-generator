package main

import (
	"fmt"
	"os"

	"github.com/koudaiii/kong-oauth-token-generator/config"
	"github.com/koudaiii/kong-oauth-token-generator/server"
)

func main() {
	printVersion()

	config, err := config.LoadConfig()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}

	server.Run(config)
}