package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yumubi/bookmarks.git/internal/api"
	"github.com/yumubi/bookmarks.git/internal/config"
	"github.com/yumubi/bookmarks.git/internal/domain"
	"log"
)

type App struct {
	Router *gin.Engine
	Cfg    config.AppConfig
}

func NewApp(cfg config.AppConfig) *App {
	logger := config.NewLogger(cfg)
	db := config.GetDb(cfg)

	repo := domain.NewBookRepository(db, logger)
	handler := api.NewBookmarkController(repo, logger)

	router := gin.Default()
	router.GET("/api/bookmarks", handler.GetAll)

	return &App{
		Cfg:    cfg,
		Router: router,
	}
}

func (app App) Run() {
	log.Fatal(app.Router.Run(fmt.Sprintf(":%d", app.Cfg.ServerPort)))
}
