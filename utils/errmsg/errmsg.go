package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// 文章模块的错误
	ERROR_ARTICLE_NOT_EXIST = 2001
	// 分类模块的错误
	ERROR_CATEGORY_NAME_USED = 3001
	ERROR_CATEGORY_NOT_EXIST = 3002

	// 参数绑定错误
	ERROR_REQUEST_BODY = 5001
)

var codeMsg = map[int]string{
	SUCCESS:              "OK",
	ERROR:                "FAIL",
	ERROR_USERNAME_USED:  "Username Exists",
	ERROR_PASSWORD_WRONG: "Password Wrong",
	ERROR_USER_NOT_EXIST: "User Not Exists",

	ERROR_TOKEN_EXIST:      "token not exists, please log in again",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_ARTICLE_NOT_EXIST: "Article Not Exists",

	ERROR_CATEGORY_NAME_USED: "Category Exists",
	ERROR_CATEGORY_NOT_EXIST: "Category Not Exists",

	ERROR_REQUEST_BODY: "Invalid request body",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
