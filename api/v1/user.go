package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	code := errmsg.SUCCESS

	// Check if username is already existing
	exist, _ := model.CheckUsername(data.Username)
	if exist == false {
		if err := model.CreateUser(&data); err != nil {
			code = errmsg.ERROR
		}
	} else {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	code := errmsg.SUCCESS

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	users, err := model.GetUsers(pageSize, pageNum)
	if err != nil {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    users,
		"count":   len(users),
		"message": errmsg.GetErrMsg(code),
	})
}

func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user model.User
	_ = c.ShouldBindJSON(&user)

	code := errmsg.SUCCESS

	exist, _ := model.CheckUsername(user.Username)
	if exist == false {
		err := model.EditUser(id, &user)
		if err != nil {
			code = errmsg.ERROR
		}
	} else {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})

}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := errmsg.SUCCESS

	err := model.DeleteUser(id)

	if err != nil {
		code = errmsg.ERROR
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
