package cmd

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yumubi/bookmarks.git/internal/api"
	"github.com/yumubi/bookmarks.git/internal/config"
	"github.com/yumubi/bookmarks.git/internal/domain"
	"log"
	"net/http"
)

type App struct {
	Router *gin.Engine
	Cfg    config.AppConfig
}

func NewApp(cfg config.AppConfig) *App {
	logger := config.NewLogger(cfg)
	// Using pgx driver
	db := config.GetDb(cfg, logger)
	repo := domain.NewBookmarkRepository(db, logger)

	// Using GORM
	//gormDb := config.GetGormDb(cfg, logger)
	//repo := domain.NewGormBookmarkRepository(gormDb, logger)

	bookmarkController := api.NewBookmarkController(repo, logger)

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	router.Use(cors.New(corsConfig))

	router.GET("/health", healthCheckHandler)

	router.GET("/api/bookmarks", bookmarkController.GetAll)
	router.GET("/api/bookmarks/:id", bookmarkController.GetById)
	router.POST("/api/bookmarks", bookmarkController.Create)
	router.PUT("/api/bookmarks/:id", bookmarkController.Update)
	router.DELETE("/api/bookmarks/:id", bookmarkController.Delete)

	return &App{
		Cfg:    cfg,
		Router: router,
	}
}

func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
func (app App) Run() {
	log.Fatal(app.Router.Run(fmt.Sprintf(":%d", app.Cfg.ServerPort)))
}
