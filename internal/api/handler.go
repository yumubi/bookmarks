package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yumubi/bookmarks.git/internal/config"
	"github.com/yumubi/bookmarks.git/internal/domain"
	"net/http"
)

type BookmarkController struct {
	repo   domain.BookmarkRepository
	logger *config.Logger
}

func NewBookmarkController(repo domain.BookmarkRepository, logger *config.Logger) BookmarkController {
	return BookmarkController{
		repo:   repo,
		logger: logger,
	}
}

func (p BookmarkController) GetAll(c *gin.Context) {
	p.logger.Info("Finding all bookmarks")
	ctx := c.Request.Context()
	bookmarks, err := p.repo.GetAll(ctx)
	if err != nil {
		if err != nil {
			p.logger.Errorf("Error :%v", err)
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch bookmarks",
		})
		return
	}
	c.JSON(http.StatusOK, bookmarks)

}
