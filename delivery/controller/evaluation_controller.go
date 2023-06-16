package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/delivery/api/response"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type EvaluationController struct {
	router *gin.Engine
	uc     usecase.EvaluationUsecase
	user   usecase.UserUsecase
	auth   gin.IRoutes
	api.BaseApi
}

func (p *EvaluationController) listHandler(c *gin.Context) {
	payload, err := p.uc.FindAll()
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func (p *EvaluationController) getHandler(c *gin.Context) {
	id := c.Param("id")
	payload, err := p.uc.FindById(id)
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func (p *EvaluationController) getByProgramHandler(c *gin.Context) {
	var payload response.EvaluationResponse
	id := c.Param("id")
	email, _ := c.Get("Email")
	user, err := p.user.SearchEmail(email.(string))
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	payload.Mid, err = p.uc.GetByProgramUser(id, user.ID, "mid")
	payload.Final, err = p.uc.GetByProgramUser(id, user.ID, "final")
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func (p *EvaluationController) createHandler(c *gin.Context) {
	var evaluation model.Evaluation

	if err := p.ParseRequestBody(c, &evaluation); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	evaluation.IsEvaluated = false

	// create save for stage mid -> half
	evaluation.Stage = "mid"
	if err := p.uc.SaveData(&evaluation); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	evaluation.ID = ""
	evaluation.Stage = "final"
	if err := p.uc.SaveData(&evaluation); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, evaluation, "OK")
}

func (p *EvaluationController) updateHandler(c *gin.Context) {
	var evaluation model.Evaluation

	if err := p.ParseRequestBody(c, &evaluation); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, "Question Category not valid")
		return
	}
	if err := p.uc.SaveData(&evaluation); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, evaluation, "OK")
}

func (p *EvaluationController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := p.uc.DeleteData(id); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func (p *EvaluationController) getEvaluationHandler(c *gin.Context) {
	panelistId := c.Param("id")
	programId := c.Param("programId")
	payload, err := p.uc.GetEvaluateeByProgramPanelist(programId, panelistId)
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func NewEvaluationController(r *gin.Engine, auth gin.IRoutes, uc usecase.EvaluationUsecase, user usecase.UserUsecase) *EvaluationController {
	controller := EvaluationController{
		router: r,
		uc:     uc,
		user:   user,
		auth:   auth,
	}
	auth.GET("/evaluation/program/:id", controller.getByProgramHandler)
	r.GET("/evaluation", controller.listHandler)
	r.GET("/evaluation/:id", controller.getHandler)
	r.GET("/evaluation/:id/:programId", controller.getEvaluationHandler)
	r.POST("/evaluation", controller.createHandler)
	r.PUT("/evaluation", controller.updateHandler)
	r.DELETE("/evaluation/:id", controller.deleteHandler)
	return &controller
}
