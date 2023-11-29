package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddArticle(c *gin.Context) {
	var article model.Article
	_ = c.ShouldBindJSON(&article)
	code := errmsg.SUCCESS

	err := model.CreateArticle(&article)

	if err != nil {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticlesByCategoryId(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	code := errmsg.SUCCESS

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, err := model.GetArticlesByCategoryId(id, pageSize, pageNum)
	count, _ := model.GetArticleCountByCategoryId(id)

	if err != nil {
		code = errmsg.ERROR_CATEGORY_NOT_EXIST
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   count,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := errmsg.SUCCESS

	data, err := model.GetArticleById(id)

	if err != nil {
		code = errmsg.ERROR_ARTICLE_NOT_EXIST
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArticles will return all articles
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")
	code := errmsg.SUCCESS

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	if len(title) == 0 {
		data, err := model.GetArticles(pageSize, pageNum)
		count, _ := model.GetArticleCount()

		if err != nil {
			code = errmsg.ERROR
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   count,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	data, err := model.SearchArticlesByTitle(title, pageSize, pageNum)
	count, _ := model.GetArticleCountByTitle(title)

	if err != nil {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   count,
		"message": errmsg.GetErrMsg(code),
	})
}

func EditArticle(c *gin.Context) {
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := errmsg.SUCCESS

	err := model.EditArticle(id, &data)

	if err != nil {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := errmsg.SUCCESS

	err := model.DeleteArticle(id)

	if err != nil {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
