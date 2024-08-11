package middleware

import (
	"api_gateway/pkg/jwt"
	"fmt"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type casbinPermission struct {
	enforcer *casbin.Enforcer
}

func (c *casbinPermission) getRole(ctx *gin.Context) (string, error) {

	token, err := ctx.Cookie("access_token")
	if err != nil {
		return "", fmt.Errorf("error while taking access token from cookie: %s", err.Error())
	}

	user, err := jwt.ExtractClaims(token, false)
	if err != nil {
		return "", fmt.Errorf("error while extracting token claims: %s", err.Error())
	}

	return user["user_role"].(string), nil
}

func (c *casbinPermission) CheckPermission(ctx *gin.Context) (bool, error) {

	var (
		sub, obj, act string
		err           error
	)

	act = ctx.Request.Method
	sub, err = c.getRole(ctx)
	if err != nil {
		return false, err
	}
	obj = ctx.Request.URL.String()

	access, err := c.enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, err
	}

	return access, nil
}

func CheckPermissionMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	casbHandler := casbinPermission{
		enforcer: enforcer,
	}

	return func(ctx *gin.Context) {

		access, err := casbHandler.CheckPermission(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
			return
		}

		if !access {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "user do not have right access",
				"error":   "Unauthorized",
			})
			return
		}

		ctx.Next()
	}
}

func MiddlewareJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "error while getting access toke from cookie",
				"error":   err.Error(),
			})
			return
		}

		valid, err := jwt.ValidateToken(token, false)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "error while extracting token claims",
				"error":   err.Error(),
			})
			return
		}

		if !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "token is expired, you should refresh token",
			})
			return
		}

		ctx.Next()
	}
}
