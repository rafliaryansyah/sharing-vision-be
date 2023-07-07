package controller

import (
	"ArticleSV/models"
	"ArticleSV/service"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"math"
	"net/http"
	"strconv"
)

func CreateArticle(c *gin.Context) {
	var Article models.Article
	if err := c.BindJSON(&Article); err != nil {
		log.Fatal(err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(Article); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]models.ApiError, len(ve))
			for i, fe := range ve {
				out[i] = models.ApiError{fe.Field(), models.MsgForTag(fe)}
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	_, err := service.AddArticle(Article)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve articles",
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    200,
		"status":  "OK",
		"message": "Berhasil membuat artikel baru.",
	})
}

func RetrieveArticles(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid limit",
			"message": err.Error(),
		})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid offset",
			"message": err.Error(),
		})
		return
	}

	data, err := service.GetArticles(limit, offset)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve articles",
			"message": err.Error(),
		})
		return
	}

	total, err := service.GetTotalArticles()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve articles",
			"message": err.Error(),
		})
		return
	}

	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": "OK",
		"code":   200,
		"data":   data,
		"meta": gin.H{
			"page":      offset,
			"perPage":   limit,
			"total":     total,
			"totalPage": totalPage,
		},
	})
}

func RetrieveArticle(c *gin.Context) {
	id := c.Param("id")
	data, err := service.GetArticleByIdOrSlug(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":   200,
		"status": "OK",
		"data":   data,
	})
}

func UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	_, err := service.GetArticleByIdOrSlug(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	var Article models.Article
	log.Println(Article)
	if err := c.BindJSON(&Article); err != nil {
		log.Fatal(err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(Article); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]models.ApiError, len(ve))
			for i, fe := range ve {
				out[i] = models.ApiError{fe.Field(), models.MsgForTag(fe)}
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}
	_, err = service.UpdateArticle(id, Article)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    200,
		"status":  "OK",
		"message": "Berhasil memperbarui artikel",
	})
}

func RemoveArticle(c *gin.Context) {
	_, err := service.DeleteArticle(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"code":    200,
		"status":  "OK",
		"message": "RemoveArticle",
		"id":      c.Param("id"),
	})
}
