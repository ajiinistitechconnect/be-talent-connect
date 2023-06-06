package controller

import (
	"fmt"
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/delivery/api/request"
	"github.com/alwinihza/talent-connect-be/delivery/api/response"
	"github.com/alwinihza/talent-connect-be/delivery/middleware"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/alwinihza/talent-connect-be/utils/authenticator"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router       *gin.Engine
	uc           usecase.AuthUsecase
	tokenService authenticator.AccessToken
	api.BaseApi
}

func (a *AuthController) login(c *gin.Context) {
	var payload model.UserCredentials
	var roles []string

	if err := a.ParseRequestBody(c, &payload); err != nil {
		a.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := a.uc.Login(payload)
	if err != nil {
		a.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	for _, v := range user.Roles {
		roles = append(roles, v.Name)
	}
	cred := model.TokenModel{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      roles,
	}
	tokenDetail, err := a.tokenService.CreateAccessToken(&cred)
	if err != nil {
		a.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if err := a.tokenService.StoreAccessToken(user.Email, tokenDetail); err != nil {
		a.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	// redis add token
	response := response.LoginResponse{
		AccessToken: tokenDetail.AccessToken,
		TokenModel:  cred,
	}
	a.NewSuccessSingleResponse(c, response, "OK")
}

func (a *AuthController) getRoles(c *gin.Context) {
	var roles []string

	token, err := authenticator.BindAuthHeader(c)
	if err != nil {
		c.AbortWithStatus(401)
	}

	accountDetail, err := a.tokenService.VerifyAccessToken(token)
	user, err := a.uc.GetUserByEmail(accountDetail.Email)
	if err != nil {
		c.AbortWithStatus(401)
	}
	for _, v := range user.Roles {
		roles = append(roles, v.Name)
	}
	fmt.Println(roles)
	a.NewSuccessSingleResponse(c, roles, "OK")
}

func (a *AuthController) logout(c *gin.Context) {

	token, err := authenticator.BindAuthHeader(c)
	if err != nil {
		c.AbortWithStatus(401)
	}

	accountDetail, err := a.tokenService.VerifyAccessToken(token)
	if err != nil {
		a.NewFailedResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if err = a.tokenService.DeleteAccessToken(accountDetail.AccessUUID); err != nil {
		a.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": "Success Logout",
	})
}

func (a *AuthController) forgetPassword(c *gin.Context) {
	// create token expire
}

func (a *AuthController) changePassword(c *gin.Context) {
	var req request.ChangePassword
	token, err := authenticator.BindAuthHeader(c)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	accessToken, err := a.tokenService.VerifyAccessToken(token)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	email, err := a.tokenService.FetchAccessToken(accessToken.AccessUUID)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	if err := a.ParseRequestBody(c, &req); err != nil {
		a.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := a.uc.ChangePassword(email, req); err != nil {
		a.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	a.NewSuccessSingleResponse(c, "OK", "OK")
}

func NewAuthController(r *gin.Engine, uc usecase.AuthUsecase, tokenService authenticator.AccessToken) *AuthController {
	controller := AuthController{
		router:       r,
		uc:           uc,
		tokenService: tokenService,
	}

	auth := r.Group("/auth").Use(middleware.NewTokenValidator(tokenService).RequireToken())
	r.POST("/login", controller.login)
	r.POST("/forget-password", controller.forgetPassword)
	auth.POST("/change-password", controller.changePassword)
	auth.POST("/logout", controller.logout)
	auth.GET("/roles", controller.getRoles)
	return &controller
}
