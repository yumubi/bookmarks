package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/yumubi/bookmarks.git/internal/config"
	"github.com/yumubi/bookmarks.git/internal/domain"
	"net/http"
	"strconv"
	"time"
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

// GetAll @Summary 获取所有书签
// @host      localhost:63342
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/bookmarks [get]
func (p BookmarkController) GetAll(c *gin.Context) {
	p.logger.Info("Finding all bookmarks")
	ctx := c.Request.Context()
	bookmarks, err := p.repo.GetAll(ctx)
	if err != nil {
		p.respondWithError(c, http.StatusInternalServerError, err, "Unable to fetch bookmarks")
		return
	}
	c.JSON(http.StatusOK, bookmarks)

}

func (p *BookmarkController) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.respondWithError(c, http.StatusBadRequest, err, "Invalid bookmark id")
		return
	}
	p.logger.Infof("Finding bookmark by id: %d", id)
	bookmark, err := p.repo.GetByID(ctx, id)
	if err != nil {
		p.respondWithError(c, http.StatusInternalServerError, err, "Unable to fetch bookmark by id")
		return
	}
	if bookmark == nil {
		p.respondWithError(c, http.StatusNotFound, nil, "Bookmark not found")
		return
	}
	c.JSON(http.StatusOK, bookmark)
}

func (p *BookmarkController) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var model domain.Bookmark
	if err := c.ShouldBindJSON(&model); err != nil {
		// you can extract error details as follows
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field())
			fmt.Println(err.Tag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
		}
		p.respondWithError(c, http.StatusBadRequest, err, "Invalid request payload")
		return
	}

	p.logger.Infof("Creating bookmark for URL: %s", model.Url)
	model.ID = 0
	model.CreatedAt = time.Now()

	savedBookmark, err := p.repo.Create(ctx, model)
	if err != nil {
		p.respondWithError(c, http.StatusInternalServerError, err, "Failed to create bookmark")
		return
	}
	c.JSON(http.StatusCreated, savedBookmark)
}

func (p *BookmarkController) Update(c *gin.Context) {
	ctx := c.Request.Context()
	var model domain.Bookmark
	if err := c.ShouldBindJSON(&model); err != nil {
		p.respondWithError(c, http.StatusBadRequest, err, "Invalid request payload")
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.respondWithError(c, http.StatusBadRequest, err, "Invalid bookmark id")
		return
	}
	p.logger.Infof("Updating bookmark for ID: %d", model.ID)
	model.ID = id
	err = p.repo.Update(ctx, model)
	if err != nil {
		p.respondWithError(c, http.StatusInternalServerError, err, "Failed to update bookmark")
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (p *BookmarkController) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		p.respondWithError(c, http.StatusBadRequest, err, "Invalid bookmark id")
		return
	}
	p.logger.Infof("Deleting bookmark for ID: %d", id)
	err = p.repo.Delete(ctx, id)
	if err != nil {
		p.respondWithError(c, http.StatusInternalServerError, err, "Failed to delete bookmark")
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (p *BookmarkController) respondWithError(c *gin.Context, code int, err error, errMsg string) {
	if err != nil {
		p.logger.Errorf("Error :%v", err)
	}
	c.AbortWithStatusJSON(code, gin.H{
		"error": errMsg,
	})
}
