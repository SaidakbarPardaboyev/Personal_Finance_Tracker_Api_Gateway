package api

import (
	_ "api_gateway/api/docs"
	v1 "api_gateway/api/handlers/v1"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/storage"

	"github.com/gin-gonic/gin"
	swaggerFile "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
)

// @title LinguaLeap
// @version 1.0
// @description Something big

// @contact.url http://www.support_me_with_smile

// @BasePath /
func NewRouter(log logger.ILogger, services client.IServiceManager, storage storage.IStorage) *gin.Engine {
	handlerV1 := v1.NewHandlerV1(services, storage, log)

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
	{
		users.GET("/profile", handlerV1.GetUserProfile)
		users.GET("/:user_id", handlerV1.GetUserById)
		users.GET("/all", handlerV1.GetAllUsers)

		users.PUT("/update", handlerV1.UpdateUserProfile)
		users.PUT("/password", handlerV1.ChangePassword)
	}

	return r
}
