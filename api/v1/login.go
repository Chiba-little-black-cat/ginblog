package v1

import (
	"errors"
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var user model.User
	var token string
	code := errmsg.SUCCESS

	if err := c.ShouldBindJSON(&user); err != nil {
		code = errmsg.ERROR_REQUEST_BODY
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
		return
	}

	isValid, err := validatePassword(user.Username, user.Password)

	switch {
	case isValid:
		if token, err = middleware.GenerateJWT(user.Username); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  errmsg.ERROR,
				"message": "Failed to generate JWT",
			})
		}
	case errors.Is(err, model.ErrUserNotFound):
		code = errmsg.ERROR_USER_NOT_EXIST
	case errors.Is(err, utils.ErrWrongPassword):
		code = errmsg.ERROR_PASSWORD_WRONG
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    user.Username,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

func validatePassword(username, password string) (bool, error) {
	encryptedPassword, err := model.GetPasswordByUsername(username)
	if err != nil {
		return false, err
	}

	isValid, err := utils.ValidatePassword(password, encryptedPassword)

	return isValid, err
}
