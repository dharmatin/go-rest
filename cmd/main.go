package main

import (
	"flag"
	"os"

	server "../app/server"
)

func main() {
	configPath := flag.String("configpath", "config.json", "Configuration path")
	flag.Parse()
	os.Setenv("CONFIG_PATH", *configPath)
	server.Start()
}
