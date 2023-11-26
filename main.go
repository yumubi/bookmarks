package main

import (
	"flag"
	"github.com/yumubi/bookmarks.git/cmd"
	"github.com/yumubi/bookmarks.git/internal/config"
	"log"
)

func main() {
	cfg, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	app := cmd.NewApp(cfg)
	app.Run()
}

func getConfig() config.AppConfig {
	var confFile string
	flag.StringVar(&confFile, "conf", "config.json", "config path, eg: -conf dev.json")
	flag.Parse()
	cfg, err := config.GetConfig(confFile)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
