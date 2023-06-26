package controller

import (
	"fmt"
	"net/http"

	"github.com/alwinihza/talent-connect-be/config"
	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/delivery/api/request"
	"github.com/alwinihza/talent-connect-be/delivery/api/response"
	"github.com/alwinihza/talent-connect-be/delivery/middleware"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/alwinihza/talent-connect-be/utils"
	"github.com/alwinihza/talent-connect-be/utils/authenticator"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router       *gin.Engine
	uc           usecase.AuthUsecase
	tokenService authenticator.AccessToken
	cfg          config.Config
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
		ID:        user.ID,
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

func (a *AuthController) redirectGoogle(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func (a *AuthController) loginGoogle(ctx *gin.Context) {
	var roles []string

	code := ctx.Query("code")

	// if ctx.Query("state") != "" {
	// 	pathUrl = ctx.Query("state")
	// }

	if code == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
		return
	}

	tokenRes, err := utils.GetGoogleOauthToken(a.cfg, code)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	fmt.Println("token", tokenRes)

	userGoogle, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)
	fmt.Println("User", userGoogle)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	user, err := a.uc.GetUserByEmail(userGoogle.Email)

	if err != nil {
		a.NewFailedResponse(ctx, http.StatusInternalServerError, err.Error())
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
		ID:        user.ID,
	}
	tokenDetail, err := a.tokenService.CreateAccessToken(&cred)
	if err != nil {
		a.NewFailedResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if err := a.tokenService.StoreAccessToken(user.Email, tokenDetail); err != nil {
		a.NewFailedResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	// redis add token
	response := response.LoginResponse{
		AccessToken: tokenDetail.AccessToken,
		TokenModel:  cred,
	}
	// set cookie and redirect to oauth url
	a.NewSuccessSingleResponse(ctx, response, "OK")
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

func NewAuthController(r *gin.Engine, uc usecase.AuthUsecase, tokenService authenticator.AccessToken, cfg config.Config) *AuthController {
	controller := AuthController{
		router:       r,
		uc:           uc,
		tokenService: tokenService,
		cfg:          cfg,
	}

	auth := r.Group("/auth").Use(middleware.NewTokenValidator(tokenService).RequireToken())
	r.POST("/login", controller.login)
	r.POST("/forget-password", controller.forgetPassword)
	r.GET("/sessions/oauth/google", controller.redirectGoogle)
	r.GET("/sessions/oauth", controller.loginGoogle)
	auth.POST("/change-password", controller.changePassword)
	auth.POST("/logout", controller.logout)
	auth.GET("/roles", controller.getRoles)
	return &controller
}
