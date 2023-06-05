package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type EvaluationCategoryController struct {
	router *gin.Engine
	uc     usecase.EvaluationCategoryUsecase
	api.BaseApi
}

func (p *EvaluationCategoryController) listHandler(c *gin.Context) {
	payload, err := p.uc.FindAll()
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func (p *EvaluationCategoryController) getHandler(c *gin.Context) {
	id := c.Param("id")
	payload, err := p.uc.FindById(id)
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func (p *EvaluationCategoryController) createHandler(c *gin.Context) {
	var evaluation_category model.EvaluationCategoryQuestion

	if err := p.ParseRequestBody(c, &evaluation_category); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := p.uc.SaveData(&evaluation_category); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, evaluation_category, "OK")
}

func (p *EvaluationCategoryController) updateHandler(c *gin.Context) {
	var evaluation_category model.EvaluationCategoryQuestion

	if err := p.ParseRequestBody(c, &evaluation_category); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, "Question Category not valid")
		return
	}
	if err := p.uc.SaveData(&evaluation_category); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, evaluation_category, "OK")
}

func (p *EvaluationCategoryController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := p.uc.DeleteData(id); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewEvaluationCategoryController(r *gin.Engine, uc usecase.EvaluationCategoryUsecase) *EvaluationCategoryController {
	controller := EvaluationCategoryController{
		router: r,
		uc:     uc,
	}
	r.GET("/category/evaluation", controller.listHandler)
	r.GET("/category/evaluation/:id", controller.getHandler)
	r.POST("/category/evaluation", controller.createHandler)
	r.PUT("/category/evaluation", controller.updateHandler)
	r.DELETE("/category/evaluation/:id", controller.deleteHandler)
	return &controller
}
