package article

import (
	"github.com/gin-gonic/gin"
	"go-articles/model"
	"go-articles/pkg/log"
	"go-articles/store"
	"net/http"
)

type ArticleCreateRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Text string `json:"text" binding:"required"`
}

func HandleCreateArticle(c *gin.Context, st *store.Store, logger log.Logger) {
	var req ArticleCreateRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	a := model.Article{
		State: "NEW",
		Title: req.Title,
		Text:  req.Text,
	}
	logger.Infof("Created article: %s", a)

	uid, err := st.Articles().Create(&a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error creating article": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "Article created successfully",
		"userId": uid,
	})
}

func HandleGetArticle(c *gin.Context, st *store.Store, logger log.Logger) {
	uid := c.Param("articleID")

	u, err := st.Articles().FindById(uid)
	if err != nil {
		if err == store.ErrArticlesNotFound {
			c.JSON(http.StatusNotFound, gin.H{"User does not exist": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{
		"User": u,
	})
}

func HandleListArticle(c *gin.Context, st *store.Store, logger log.Logger) {
	u, err := st.Articles().ListArticles()
	if err != nil {
		if err == store.ErrArticlesNotFound {
			c.JSON(http.StatusNotFound, gin.H{"User does not exist": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{
		"User": u,
	})
}