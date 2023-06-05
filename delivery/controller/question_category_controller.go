package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type QuestionCategoryController struct {
	router *gin.Engine
	uc     usecase.QuestionCategoryUsecase
	api.BaseApi
}

func (p *QuestionCategoryController) listHandler(c *gin.Context) {
	question_categorys, err := p.uc.FindAll()
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, question_categorys, "OK")
}

func (p *QuestionCategoryController) getHandler(c *gin.Context) {
	id := c.Param("id")
	payload, err := p.uc.FindById(id)
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func (p *QuestionCategoryController) createHandler(c *gin.Context) {
	var question_category model.QuestionCategory

	if err := p.ParseRequestBody(c, &question_category); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := p.uc.SaveData(&question_category); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, question_category, "OK")
}

func (p *QuestionCategoryController) updateHandler(c *gin.Context) {
	var question_category model.QuestionCategory

	if err := p.ParseRequestBody(c, &question_category); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, "Question Category not valid")
		return
	}
	if err := p.uc.SaveData(&question_category); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, question_category, "OK")
}

func (p *QuestionCategoryController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := p.uc.DeleteData(id); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewQuestionCategoryController(r *gin.Engine, uc usecase.QuestionCategoryUsecase) *QuestionCategoryController {
	controller := QuestionCategoryController{
		router: r,
		uc:     uc,
	}
	r.GET("/category/question", controller.listHandler)
	r.GET("/category/question/:id", controller.getHandler)
	r.POST("/category/question", controller.createHandler)
	r.PUT("/category/question", controller.updateHandler)
	r.DELETE("/category/question/:id", controller.deleteHandler)
	return &controller
}
