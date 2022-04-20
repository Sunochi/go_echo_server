package main

import (
	"flag"
	"go_example_server/config"
	"go_example_server/database"
	"go_example_server/server"
)

var configFile = flag.String("config", "config.json", "Location of the config file.")
var logDirectory = flag.String("log", "log", "Directory path of the log file.")
var port = flag.String("port", "8080", "Run on this port.")

func main() {
	flag.Parse()
	config.Init(configFile)
	database.Init()

	router := server.Init(logDirectory)
	err := router.Start(":" + *port)
	if err != nil {
		panic(err)
	}
}
