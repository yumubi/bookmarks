package main

import (
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
