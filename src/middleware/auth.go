package middleware

import (
	"context"
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vkevv/go-graphql/src/config"
)

type userKeyType string

const userKeyContext userKeyType = "userId"
const headerAuth string = "Authorization"

// UserKey is the key where userID is stored in map claims
const UserKey string = "userID"

func tokenWithoutBearer(token string) string {
	bearer := "BEARER"
	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:]
	}
	return ""
}

// AuthMiddleware will store userID in context to use in resolvers
// if any error occurs it just will continue with the next handler
func AuthMiddleware(conf config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := tokenWithoutBearer(c.GetHeader(headerAuth))
		if tokenStr == "" {
			c.Next()
			return
		}
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(conf.App.JWT), nil
		})
		if err != nil {
			c.Next()
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.Next()
			return
		}
		userID, exist := claims[UserKey]
		if !exist {
			c.Next()
			return
		}
		ctx := context.WithValue(c.Request.Context(), userKeyContext, userID)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GetUserIDFromCtx receive a context and returns the userID if its stored
func GetUserIDFromCtx(ctx context.Context) (string, error) {
	errorNoUser := errors.New("No userID found in context")
	if ctx.Value(userKeyContext) == nil {
		return "", errorNoUser
	}
	userID, ok := ctx.Value(userKeyContext).(string)
	if !ok {
		return "", errorNoUser
	}
	return userID, nil
}
