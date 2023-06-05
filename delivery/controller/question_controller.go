package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type QuestionController struct {
	router *gin.Engine
	uc     usecase.QuestionUsecase
	api.BaseApi
}

func (p *QuestionController) listHandler(c *gin.Context) {
	questions, err := p.uc.FindAll()
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, questions, "OK")
}

func (p *QuestionController) getHandler(c *gin.Context) {
	id := c.Param("id")
	payload, err := p.uc.FindById(id)
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func (p *QuestionController) createHandler(c *gin.Context) {
	var question model.Question

	if err := p.ParseRequestBody(c, &question); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := p.uc.SaveData(&question); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, question, "OK")
}

func (p *QuestionController) updateHandler(c *gin.Context) {
	var question model.Question

	if err := p.ParseRequestBody(c, &question); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, "Question not valid")
		return
	}
	if err := p.uc.SaveData(&question); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, question, "OK")
}

func (p *QuestionController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := p.uc.DeleteData(id); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewQuestionController(r *gin.Engine, uc usecase.QuestionUsecase) *QuestionController {
	controller := QuestionController{
		router: r,
		uc:     uc,
	}
	r.GET("/questions", controller.listHandler)
	r.GET("/questions/:id", controller.getHandler)
	r.POST("/questions", controller.createHandler)
	r.PUT("/questions", controller.updateHandler)
	r.DELETE("/questions/:id", controller.deleteHandler)
	return &controller
}
