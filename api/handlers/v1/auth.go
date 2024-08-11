package v1

import (
	"api_gateway/api/handlers/models"
	pb "api_gateway/genproto/users"
	"api_gateway/pkg/helper"
	"api_gateway/pkg/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

// Register 		godoc
// @Router 			/auth/register [post]
// @Summery 		Register User
// @Description 	this API for register new user
// @Tags			Auth
// @Accept 			json
// @Produce 		json
// @Param 			register body models.RequestRegister true "register"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) Register(ctx *gin.Context) {
	var req pb.CreateUser

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		handleResponse(ctx, h.log, "Error with decoding url body", http.StatusBadRequest, err.Error())
		return
	}

	// check email is valid
	err = helper.CheckEmailAndPasswordValid(req.Email, req.Password)
	if err != nil {
		handleResponse(ctx, h.log, "Error while checking email and password is valid and strong", http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		handleResponse(ctx, h.log, "Error with hashing password", http.StatusInternalServerError, err.Error())
		return
	}
	req.Password = string(hashedPassword)

	resp, err := h.services.AuthService().Create(ctx, &req)
	if err != nil {
		handleResponse(ctx, h.log, "Error with register User", http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

// Login 			godoc
// @Router 			/auth/login [post]
// @Summery 		Login User
// @Description 	this API for login new user
// @Tags			Auth
// @Accept 			json
// @Produce 		json
// @Param 			login body models.RequestLogin true "login"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) Login(ctx *gin.Context) {
	req := models.RequestLogin{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		handleResponse(ctx, h.log, "Error with decoding url body", http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.AuthService().GetByEmail(ctx, &pb.Email{Email: req.Email})
	if err != nil {
		handleResponse(ctx, h.log, "error getting email from database", http.StatusBadRequest, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		handleResponse(ctx, h.log, "Password is incorrect", http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := jwt.GenarateJWTToken(user)
	if err != nil {
		handleResponse(ctx, h.log, "Error while generating tokens", http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.services.AuthService().DeleteRefreshTokenByUserId(ctx, &pb.PrimaryKey{Id: user.Id})
	if err != nil {
		handleResponse(ctx, h.log, "Error with deleting userinfo from refreshToken table", http.StatusInternalServerError, err.Error())
		return
	}

	_, err = h.services.AuthService().StoreRefreshToken(ctx, &pb.RefreshToken{
		UserId:       user.Id,
		RefreshToken: tokens.RefreshToken,
		ExpiresIn:    time.Now().Add(time.Hour * 24).Format(time.RFC3339),
	})
	if err != nil {
		handleResponse(ctx, h.log, "Error with storing refresh token to refreshToken table", http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("access_token", tokens.AccessToken, int(time.Hour), "/", "", false, false)
	ctx.SetCookie("refresh_token", tokens.RefreshToken, int(time.Hour*24), "/", "", false, false)

	ctx.JSON(http.StatusOK, tokens)
}

// RefreshToken 	godoc
// @Router 			/auth/refresh-token [post]
// @Summery 		refresh user access token
// @Description 	this API for refresh user access token
// @Tags 			Auth
// @Accept 			json
// @Produce 		json
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) RefreshToken(ctx *gin.Context) {

	token, err := ctx.Cookie("refresh_token")
	if err != nil {
		handleResponse(ctx, h.log, "error while getting refresh toke from cookie", http.StatusBadRequest, err.Error())
		return
	}
	req := pb.RequestRefreshToken{
		RefreshToken: token,
	}

	_, err = h.services.AuthService().CheckRefreshTokenExists(ctx, &req)
	if err != nil {
		handleResponse(ctx, h.log, "Invalid refresh token", http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := jwt.GenarateAccessToken(req.RefreshToken)
	if err != nil {
		handleResponse(ctx, h.log, "Error with generate access token", http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("access_token", tokens.AccessToken, int(time.Hour), "/", "", false, false)

	ctx.JSON(http.StatusOK, tokens)
}

// ForgotPassword 	godoc
// @Summery 		Forgot-Password
// @Router 			/auth/forgot-password [post]
// @Description 	it is used when user forgot password
// @Tags 			Auth
// @Accept 			json
// @Produce 		json
// @Param 			forgot_password body models.ForgotPasswordReq true "forgot_password"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) ForgotPassword(ctx *gin.Context) {

	var (
		request = pb.Email{}
		code    string
		err     error
	)

	if err = json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		handleResponse(ctx, h.log, "Error with decoding url body", http.StatusBadRequest, err.Error())
		return
	}

	// check email is exists
	if _, err = h.services.AuthService().CheckEmailExists(ctx, &request); err != nil {
		handleResponse(ctx, h.log, "Error with checking email is exists", http.StatusBadRequest, err.Error())
		return
	}

	code = helper.RandomCodeMaker()

	err = jwt.SendEmail(request.Email, "Reset-Password Code", "Your verification code : "+code)
	if err != nil {
		handleResponse(ctx, h.log, "error while sending message to email in service layer", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.RedisClient().SaveCodeWithEmail(ctx, request.Email, code); err != nil {
		handleResponse(ctx, h.log, "error while saving message in redis in service layer", http.StatusBadRequest, err.Error())
		return
	}

	ctx.SetCookie("email", request.Email, int(time.Minute*2), "/", "", false, false)

	handleResponse(ctx, h.log, "", http.StatusOK, "Password reset code sent to your email")
}

// ResetPassword 	godoc
// @Router 			/auth/reset-password [post]
// @Summery 		reset user password
// @Description 	this API for reset user password
// @Tags 			Auth
// @Accept 			json
// @Produce			json
// @Param 			reset-password body models.ResetPasswordReq true "reset-password"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
func (h *HandlerV1) ResetPassword(ctx *gin.Context) {

	var (
		request *models.ResetPasswordReq
		code    string
		err     error
		email   string
	)

	if err = json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		handleResponse(ctx, h.log, "Error with decoding url body", http.StatusBadRequest, err.Error())
		return
	}

	email, err = ctx.Cookie("email")
	if err != nil {
		handleResponse(ctx, h.log, "Error with taking email from cookie", http.StatusBadRequest, err.Error())
		return
	}

	code, err = h.storage.RedisClient().GetCodeWithEmail(ctx, email)
	if err != nil {
		handleResponse(ctx, h.log, "error while taking message in redis", http.StatusBadRequest, fmt.Sprintf("code was expired %s", err))
		return
	}

	if request.VerificationCode != code {
		handleResponse(ctx, h.log, "code is not correct", http.StatusBadRequest, "code is not correct")
		return
	}

	// check email is valid
	err = helper.CheckPasswordIsStrong(request.NewPassword)
	if err != nil {
		handleResponse(ctx, h.log, "Error while checking email and password is valid and strong", http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		handleResponse(ctx, h.log, "Error while hashing new password", http.StatusInternalServerError, err.Error())
		return
	}
	request.NewPassword = string(hashedPassword)

	if _, err = h.services.AuthService().ResetPassword(ctx, &pb.ResetPassword{
		NewPassword: request.NewPassword,
		Email:       email,
	}); err != nil {
		handleResponse(ctx, h.log, "error while saving new password in service layer", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "", http.StatusOK, "Password successfully reset")
}
