package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/delivery/api/request"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type QuestionAnswerController struct {
	router *gin.Engine
	uc     usecase.QuestionAnswerUsecase
	auth   gin.IRoutes
	api.BaseApi
}

func (p *QuestionAnswerController) getHandler(c *gin.Context) {
	id := c.Param("id")
	payload, err := p.uc.FindById(id)
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func (p *QuestionAnswerController) listHandler(c *gin.Context) {
	id := c.Param("evalId")

	payload, err := p.uc.GetByEvaluation(id)
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func (p *QuestionAnswerController) createHandler(c *gin.Context) {
	var payload request.AnswerRequest

	if err := p.ParseRequestBody(c, &payload); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := p.uc.SaveQuestionAnswer(&payload); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, payload, "OK")
}

func NewQuestionAnswerController(r *gin.Engine, auth gin.IRoutes, uc usecase.QuestionAnswerUsecase) *QuestionAnswerController {
	controller := QuestionAnswerController{
		router: r,
		uc:     uc,
		auth:   auth,
	}
	auth.GET("/answers/:evalId", controller.listHandler)
	r.GET("/answer/:id", controller.getHandler)
	r.POST("/answer", controller.createHandler)
	// r.PUT("/programs", controller.updateHandler)
	// r.DELETE("/programs/:id", controller.deleteHandler)
	return &controller
}
