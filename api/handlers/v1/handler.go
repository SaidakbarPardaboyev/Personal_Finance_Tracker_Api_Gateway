package v1

import (
	"api_gateway/api/handlers/models"
	"api_gateway/api/handlers/tokens"
	"api_gateway/grpc/client"
	"api_gateway/pkg/logger"
	"api_gateway/storage"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Option struct {
	Log      logger.ILogger
	Services client.IServiceManager
	Storage  storage.IStorage
}

type HandlerV1 struct {
	storage  storage.IStorage
	services client.IServiceManager
	log      logger.ILogger
}

func NewHandlerV1(option *Option) *HandlerV1 {
	return &HandlerV1{
		storage:  option.Storage,
		services: option.Services,
		log:      option.Log,
	}
}

func handleResponse(ctx *gin.Context, log logger.ILogger, msg string, statusCode int, data interface{}) {

	var (
		resp = models.Response{}
	)

	switch code := statusCode; {
	case code < 400:
		resp.Description = "OK"
		log.Info("~~~~> OK", logger.String("msg", msg), logger.Any("status", statusCode))
	case code == 401:
		resp.Description = "Unauthorized"
		log.Info("????? Unauthorized", logger.String("msg", msg), logger.Any("status", statusCode))
	case code < 500:
		resp.Description = "Bad Request"
		log.Info("!!!!! Bad Request", logger.String("msg", msg), logger.Any("status", statusCode), logger.Any("Error", data))
	default:
		resp.Description = "Internal Server Error"
		log.Info("!!!!! Internal Server Error", logger.String("msg", msg), logger.Any("status", statusCode), logger.Any("Error", data))
	}

	resp.StatisCode = statusCode
	resp.Data = data

	ctx.JSON(statusCode, resp)
}

func getUserInfoFromToken(ctx *gin.Context) (*models.UserInfoFromToken, error) {

	var (
		token  string
		claims jwt.MapClaims
		resp   = models.UserInfoFromToken{}
		err    error
	)

	token, err = ctx.Cookie("access_token")
	if err != nil {
		return nil, fmt.Errorf("error while getting access toke from cookie: %s", err.Error())
	}

	claims, err = tokens.ExtractClaims(token)
	if err != nil {
		return nil, err
	}

	resp.Id = claims["user_id"].(string)
	resp.Email = claims["email"].(string)
	resp.FullName = claims["full_name"].(string)
	resp.UserRole = claims["user_role"].(string)

	return &resp, nil
}
