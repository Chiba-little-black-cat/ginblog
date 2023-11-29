package middleware

import (
	"errors"
	"ginblog/config"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(config.JwtKey)

type MyClaims struct {
	Username string
	jwt.RegisteredClaims
}

var (
	TokenExpired     = errors.New("expired token, please log in again")
	TokenNotValidYet = errors.New("not valid token, please login again")
	TokenMalformed   = errors.New("malformed token, please log in again")
	TokenInvalid     = errors.New("invalid token, please log in again")
)

func GenerateJWT(username string) (string, error) {
	claims := MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "GinBlog",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	switch {
	case token.Valid:
		return token, nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		return token, TokenMalformed
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return token, TokenInvalid
	case errors.Is(err, jwt.ErrTokenExpired):
		return token, TokenExpired
	case errors.Is(err, jwt.ErrTokenNotValidYet):
		return token, TokenNotValidYet
	default:
		return token, errors.New("unexpected error")
	}
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": errmsg.ERROR,
				"error":  "Missing Authorization header",
			})
			return
		}

		// Check header format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": errmsg.ERROR,
				"error":  "Invalid Token format",
			})
			return
		}

		token, err := ValidateJWT(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  errmsg.ERROR,
				"message": err.Error(),
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": errmsg.ERROR,
				"error":  "Invalid Token Claims",
			})
			return
		}

		username, exists := claims["Username"].(string)
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": errmsg.ERROR,
				"error":  "Username not found in claims",
			})
			return
		}

		if username != "admin" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": errmsg.ERROR,
				"error":  "Not Admin User",
			})
			return
		}

		c.Next()
	}
}
