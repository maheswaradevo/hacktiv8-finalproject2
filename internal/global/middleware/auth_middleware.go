package middleware

import (
	"errors"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maheswaradevo/hacktiv8-finalproject2/internal/global/utils"
	errRes "github.com/maheswaradevo/hacktiv8-finalproject2/pkg/errors"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := verifyToken(c)
		_ = verifyToken
		if err != nil {
			errResponse := utils.NewErrorResponse(c.Writer, errRes.ErrUnauthorized)
			c.JSON(errResponse.Error.Code, errResponse)
			c.Abort()
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}

func verifyToken(c *gin.Context) (interface{}, error) {
	headerToken := c.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errors.New("sign in to proceed")
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}
