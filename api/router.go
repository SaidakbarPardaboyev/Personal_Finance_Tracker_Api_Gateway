package api

import (
	_ "api_gateway/api/docs"
	v1 "api_gateway/api/handlers/v1"
	"api_gateway/api/middleware"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/storage"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"

	swagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Log            logger.ILogger
	Services       client.IServiceManager
	Storage        storage.IStorage
	CasbinEnforcer *casbin.Enforcer
}

// @title LinguaLeap
// @version 1.0
// @description Something big

// @contact.url http://www.support_me_with_smile
// @securetyDefinitions.apikay ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /
func NewRouter(option *Option) *gin.Engine {
	handlerV1 := v1.NewHandlerV1(&v1.Option{
		Log:      option.Log,
		Services: option.Services,
		Storage:  option.Storage,
	})

	r := gin.Default()

	r.GET("swagger/*any", swagger.WrapHandler(swaggerFile.Handler))

	auth := r.Group("/auth")
	{
		auth.POST("/register", handlerV1.Register)
		auth.POST("/login", handlerV1.Login)
		auth.POST("/refresh-token", handlerV1.RefreshToken)
		auth.POST("/forgot-password", handlerV1.ForgotPassword)
		auth.POST("/reset-password", handlerV1.ResetPassword)
	}

	users := r.Group("/users")
	users.Use(middleware.MiddlewareJWT())
	users.Use(middleware.CheckPermissionMiddleware(option.CasbinEnforcer))
	{
		users.GET("/profile", handlerV1.GetUserProfile)
		users.GET("/:user_id", handlerV1.GetUserById)
		users.GET("/all", handlerV1.GetAllUsers)
		users.PUT("/update", handlerV1.UpdateUserProfile)
		users.DELETE("/:user_id", handlerV1.DeleteUser)
		users.PUT("/password", handlerV1.ChangePassword)
		users.PUT("/user_role/:user_id", handlerV1.ChangeUserRole)

	}

	return r
}
