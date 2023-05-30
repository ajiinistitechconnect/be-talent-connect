package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type ProgramController struct {
	router *gin.Engine
	uc     usecase.ProgramUsecase
	api.BaseApi
}

func (p *ProgramController) listHandler(c *gin.Context) {
	programs, err := p.uc.FindAll()
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, programs, "OK")
}

func (p *ProgramController) getHandler(c *gin.Context) {
	id := c.Param("id")
	payload, err := p.uc.FindById(id)
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload, "OK")
}

func (p *ProgramController) createHandler(c *gin.Context) {
	var program model.Program

	if err := p.ParseRequestBody(c, &program); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := p.uc.SaveData(&program); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, program, "OK")
}

func (p *ProgramController) updateHandler(c *gin.Context) {
	var program model.Program

	if err := p.ParseRequestBody(c, &program); err != nil {
		p.NewFailedResponse(c, http.StatusBadRequest, "Program not valid")
		return
	}
	if err := p.uc.SaveData(&program); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	p.NewSuccessSingleResponse(c, program, "OK")
}

func (p *ProgramController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := p.uc.DeleteData(id); err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func NewProgramController(r *gin.Engine, uc usecase.ProgramUsecase) *ProgramController {
	controller := ProgramController{
		router: r,
		uc:     uc,
	}
	r.GET("/programs", controller.listHandler)
	r.GET("/programs/:id", controller.getHandler)
	r.POST("/programs", controller.createHandler)
	r.PUT("/programs", controller.updateHandler)
	r.DELETE("/programs/:id", controller.deleteHandler)
	return &controller
}
