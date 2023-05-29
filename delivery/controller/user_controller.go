package controller

import (
	"encoding/json"
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	router *gin.Engine
	uc     usecase.UserUsecase
	api.BaseApi
}

func (u *UserController) listHandler(c *gin.Context) {

	farmers, err := u.uc.FindAll()
	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, farmers, "OK")
}

func (r *UserController) createHandler(c *gin.Context) {
	user := c.PostForm("user")
	role := c.PostForm("role")
	var payload model.User
	var roleUser []string

	if err := json.Unmarshal([]byte(user), &payload); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, "User not valid")
		return
	}
	if err := json.Unmarshal([]byte(role), &roleUser); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, "Role not valid")
		return
	}

	if err := r.uc.UpdateRole(&payload, roleUser); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := r.uc.SaveData(&payload); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r.NewSuccessSingleResponse(c, payload, "OK")
}

func (r *UserController) updateHandler(c *gin.Context) {
	user := c.PostForm("user")
	role := c.PostForm("role")
	var payload model.User
	var roleUser []string

	if err := json.Unmarshal([]byte(user), &payload); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, "User not valid")
		return
	}
	if err := json.Unmarshal([]byte(role), &roleUser); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, "Role not valid")
		return
	}

	if len(roleUser) > 0 {
		if err := r.uc.UpdateRole(&payload, roleUser); err != nil {
			r.NewFailedResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	if err := r.uc.SaveData(&payload); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r.NewSuccessSingleResponse(c, payload, "OK")
}

func (r *UserController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := r.uc.DeleteData(id); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewUserController(r *gin.Engine, uc usecase.UserUsecase) *UserController {
	controller := UserController{
		router: r,
		uc:     uc,
	}
	r.GET("/users", controller.listHandler)
	r.PUT("/users/", controller.updateHandler)
	r.POST("/users", controller.createHandler)
	// r.DELETE("/farmers/:id", controller.deleteHandler)
	return &controller
}
