package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type ParticipantController struct {
	router *gin.Engine
	uc     usecase.ParticipantUsecase
	api.BaseApi
}

func (u *ParticipantController) listHandler(c *gin.Context) {
	participants, err := u.uc.FindAll()
	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, participants, "OK")
}

func (u *ParticipantController) getEvalHandler(c *gin.Context) {
	id := c.Param("id")
	participants, err := u.uc.GetEvaluationScore(id)
	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, participants, "OK")
}

func (r *ParticipantController) createHandler(c *gin.Context) {
	var payload model.Participant
	if err := r.ParseRequestBody(c, &payload); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.uc.SaveData(&payload); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r.NewSuccessSingleResponse(c, payload, "OK")
}

func (r *ParticipantController) updateHandler(c *gin.Context) {
	var payload model.Participant
	if err := r.ParseRequestBody(c, &payload); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := r.uc.SaveData(&payload); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r.NewSuccessSingleResponse(c, payload, "OK")
}

func (r *ParticipantController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := r.uc.DeleteData(id); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewParticipantController(r *gin.Engine, uc usecase.ParticipantUsecase) *ParticipantController {
	controller := ParticipantController{
		router: r,
		uc:     uc,
	}
	r.GET("/participants", controller.listHandler)
	r.GET("/participants/evaluation/:id", controller.getEvalHandler)
	r.PUT("/participants", controller.updateHandler)
	r.POST("/participants", controller.createHandler)
	r.DELETE("/participants/:id", controller.deleteHandler)
	return &controller
}
