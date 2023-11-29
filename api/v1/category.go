package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code := errmsg.SUCCESS

	// Check if category name is already existing
	exist, _ := model.CheckCategoryName(data.Name)
	if exist == false {
		if err := model.CreateCategory(&data); err != nil {
			code = errmsg.ERROR
		}
	} else {
		code = errmsg.ERROR_CATEGORY_NAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetCategories(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	code := errmsg.SUCCESS

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	categories, err := model.GetCategories(pageSize, pageNum)
	if err != nil {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    categories,
		"count":   len(categories),
		"message": errmsg.GetErrMsg(code),
	})
}

func EditCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var category model.Category
	_ = c.ShouldBindJSON(&category)
	code := errmsg.SUCCESS

	exist, _ := model.CheckCategoryName(category.Name)
	if exist == false {
		err := model.EditCategory(id, &category)
		if err != nil {
			code = errmsg.ERROR
		}
	} else {
		code = errmsg.ERROR_CATEGORY_NAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})

}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := errmsg.SUCCESS

	err := model.DeleteCategory(id)

	if err != nil {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
