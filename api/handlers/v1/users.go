package v1

import (
	"api_gateway/api/handlers/models"
	pb "api_gateway/genproto/users"
	"api_gateway/pkg/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetUserProfile  	godoc
// @Router 			/users/profile [get]
// @Summary 		Get User Profile
// @Description 	getting user profile
// @Accept 			json
// @Produce 		json
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
// @Security		ApiKeyAuth
func (h *HandlerV1) GetUserProfile(ctx *gin.Context) {

	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		handleResponse(ctx, h.log, "error while getting user info from token", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.UsersService().GetById(ctx, &pb.PrimaryKey{Id: user.Id})
	if err != nil {
		handleResponse(ctx, h.log, "error while using GetById method of users service", http.StatusInternalServerError, logger.Error(err))
		return
	}

	handleResponse(ctx, h.log, "", http.StatusOK, resp)
}

// GetUserById  	godoc
// @Router 			/users/{user_id} [get]
// @Summary 		Get User Profile by id
// @Description 	getting user profile by user id
// @Accept 			json
// @Produce 		json
// @Param			user_id path string true "user_id"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
// @Security		ApiKeyAuth
func (h *HandlerV1) GetUserById(ctx *gin.Context) {

	userId := ctx.Param("user_id")
	if userId == "" {
		handleResponse(ctx, h.log, "error: id not found in request param", http.StatusBadRequest, logger.Error(fmt.Errorf("id not found in request param")))
		return
	}

	resp, err := h.services.UsersService().GetById(ctx, &pb.PrimaryKey{Id: userId})
	if err != nil {
		handleResponse(ctx, h.log, "error while using GetById method of users service", http.StatusInternalServerError, logger.Error(err))
		return
	}

	handleResponse(ctx, h.log, "", http.StatusOK, resp)
}

// GetAllUsers  	godoc
// @Router 			/users/all [get]
// @Summary 		Get All User
// @Description 	getting all user
// @Accept 			json
// @Produce 		json
// @Param			page 		query int 	true 	"page"
// @Param			limit 		query int 	true 	"limit"
// @Param			full_name 	query string false 	"full_name"
// @Param			email 		query string false 	"email"
// @Param			user_role 	query string false 	"user_role"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
// @Security		ApiKeyAuth
func (h *HandlerV1) GetAllUsers(ctx *gin.Context) {

	pageStr := ctx.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(ctx, h.log, "error while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(ctx, h.log, "error while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	fullname := ctx.Query("full_name")
	email := ctx.Query("email")
	userRole := ctx.Query("user_role")

	resp, err := h.services.UsersService().GetAll(ctx, &pb.GetListRequest{
		Page:     int32(page),
		Limit:    int64(limit),
		FullName: fullname,
		Email:    email,
		UserRole: userRole,
	})
	if err != nil {
		handleResponse(ctx, h.log, "error while using GetAll method of users service", http.StatusInternalServerError, logger.Error(err))
		return
	}

	handleResponse(ctx, h.log, "", http.StatusOK, resp)
}

// UpdateUserProfile  	godoc
// @Router 			/users/update [put]
// @Summary 		Update User Profile
// @Description 	updating user profile
// @Accept 			json
// @Produce 		json
// @Param 			user body models.UpdateUser true "user"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
// @Security		ApiKeyAuth
func (h *HandlerV1) UpdateUserProfile(ctx *gin.Context) {

	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		handleResponse(ctx, h.log, "error while getting user info from token", http.StatusBadRequest, err.Error())
		return
	}

	req := models.UpdateUser{}
	if err = json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		handleResponse(ctx, h.log, "error while getting request body", http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		handleResponse(ctx, h.log, "Error with hashing password", http.StatusInternalServerError, err.Error())
		return
	}
	req.Password = string(hashedPassword)

	resp, err := h.services.UsersService().Update(ctx, &pb.UpdateUser{
		Id:           user.Id,
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: req.Password,
	})
	if err != nil {
		handleResponse(ctx, h.log, "error while using Update method of users service", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "", http.StatusOK, resp)
}

// DeleteUser  	godoc
// @Router 			/users/{user_id} [delete]
// @Summary 		Delete User
// @Description 	deleting user by user id
// @Accept 			json
// @Produce 		json
// @Param			user_id path string true "user_id"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
// @Security		ApiKeyAuth
func (h *HandlerV1) DeleteUser(ctx *gin.Context) {

	userId := ctx.Param("user_id")
	if userId == "" {
		handleResponse(ctx, h.log, "error: id not found in request param", http.StatusBadRequest, logger.Error(fmt.Errorf("id not found in request param")))
		return
	}

	_, err := h.services.UsersService().Delete(ctx, &pb.PrimaryKey{Id: userId})
	if err != nil {
		handleResponse(ctx, h.log, "error while using Delete method of users service", http.StatusInternalServerError, logger.Error(err))
		return
	}

	handleResponse(ctx, h.log, "Success", http.StatusOK, "user was deleted successfully")
}

// ChangePassword  	godoc
// @Router 			/users/password [put]
// @Summary 		Change User Password
// @Description 	changing user password
// @Accept 			json
// @Produce 		json
// @Param 			change_password body models.ChangePassword true "change_password"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
// @Security		ApiKeyAuth
func (h *HandlerV1) ChangePassword(ctx *gin.Context) {

	user, err := getUserInfoFromToken(ctx)
	if err != nil {
		handleResponse(ctx, h.log, "error while getting user info from token", http.StatusBadRequest, err.Error())
		return
	}

	req := models.ChangePassword{}
	if err = json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		handleResponse(ctx, h.log, "error while getting request body", http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.services.UsersService().ChangePassword(ctx, &pb.ChangePassword{
		CurrentPassword: req.CurrentPassword,
		NewPassword:     req.NewPassword,
		UserId:          user.Id,
	})
	if err != nil {
		handleResponse(ctx, h.log, "error while using ChangePassword method of users service", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "user password was changed successfully", http.StatusOK, "user password was changed successfully")
}

// ChangeUserRole  	godoc
// @Router 			/users/user_role/{user_id} [put]
// @Summary 		Change User Role
// @Description 	changing user Role
// @Accept 			json
// @Produce 		json
// @Param 			change_user_role body models.ChangeUserRole true "change_user_role"
// @Success 		200  {object}  models.Response
// @Failure 		400  {object}  models.Response
// @Failure 		500  {object}  models.Response
// @Failure 		401  {object}  models.Response
// @Security		ApiKeyAuth
func (h *HandlerV1) ChangeUserRole(ctx *gin.Context) {

	userId := ctx.Param("user_id")
	if userId == "" {
		handleResponse(ctx, h.log, "error: id not found in request param", http.StatusBadRequest, logger.Error(fmt.Errorf("id not found in request param")))
		return
	}

	request := models.ChangeUserRole{}
	if err := json.NewDecoder(ctx.Request.Body).Decode(&request); err != nil {
		handleResponse(ctx, h.log, "error while getting request body", http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.UsersService().ChangeUserRole(ctx, &pb.ChangeUserRole{
		Id:          userId,
		NewUserRole: request.NewUserRole,
	})
	if err != nil {
		handleResponse(ctx, h.log, "error while using ChangeUserRole method of users service", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(ctx, h.log, "user role was changed successfully", http.StatusOK, "user role was changed successfully")
}
