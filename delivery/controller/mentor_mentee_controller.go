package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type MentorMenteeController struct {
	router *gin.Engine
	uc     usecase.MentorMenteeUsecase
	api.BaseApi
}

func (u *MentorMenteeController) listHandler(c *gin.Context) {
	mentorMentees, err := u.uc.FindAll()
	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, mentorMentees, "OK")
}

func (r *MentorMenteeController) createHandler(c *gin.Context) {
	var payload model.MentorMentee
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

func (r *MentorMenteeController) updateHandler(c *gin.Context) {
	var payload model.MentorMentee
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

func (r *MentorMenteeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := r.uc.DeleteData(id); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewMentorMenteeController(r *gin.Engine, uc usecase.MentorMenteeUsecase) *MentorMenteeController {
	controller := MentorMenteeController{
		router: r,
		uc:     uc,
	}
	r.GET("/mentor-mentees", controller.listHandler)
	r.PUT("/mentor-mentees", controller.updateHandler)
	r.POST("/mentor-mentees", controller.createHandler)
	r.DELETE("/mentor-mentees/:id", controller.deleteHandler)
	return &controller
}
