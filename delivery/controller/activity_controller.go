package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type ActivityController struct {
	router *gin.Engine
	uc     usecase.ActivityUsecase
	api.BaseApi
}

func (a *ActivityController) listHandler(c *gin.Context) {
	activities, err := a.uc.FindAll()
	if err != nil {
		a.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	a.NewSuccessSingleResponse(c, activities, "OK")
}

func (a *ActivityController) getHandler(c *gin.Context) {
	id := c.Param("id")
	payload, err := a.uc.FindById(id)
	if err != nil {
		a.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	a.NewSuccessSingleResponse(c, payload, "OK")
}

func (a *ActivityController) createHandler(c *gin.Context) {
	var activity model.Activity

	if err := a.ParseRequestBody(c, &activity); err != nil {
		a.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.uc.SaveData(&activity); err != nil {
		a.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	a.NewSuccessSingleResponse(c, activity, "OK")
}

func (a *ActivityController) updateHandler(c *gin.Context) {
	var activity model.Activity

	if err := a.ParseRequestBody(c, &activity); err != nil {
		a.NewFailedResponse(c, http.StatusBadRequest, "Program not valid")
		return
	}
	if err := a.uc.SaveData(&activity); err != nil {
		a.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	a.NewSuccessSingleResponse(c, activity, "OK")
}

func (a *ActivityController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := a.uc.DeleteData(id); err != nil {
		a.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewActivityController(r *gin.Engine, uc usecase.ActivityUsecase) *ActivityController {
	controller := ActivityController{
		router: r,
		uc:     uc,
	}
	r.GET("/activities", controller.listHandler)
	r.GET("/activities/:id", controller.getHandler)
	r.POST("/activities", controller.createHandler)
	r.PUT("/activities", controller.updateHandler)
	r.DELETE("/activities/:id", controller.deleteHandler)
	return &controller
}
