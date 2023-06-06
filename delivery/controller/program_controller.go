package controller

import (
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/delivery/api/response"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type ProgramController struct {
	router *gin.Engine
	auth   gin.IRoutes
	uc     usecase.ProgramUsecase
	user   usecase.UserUsecase
	api.BaseApi
}

func (p *ProgramController) listHandler(c *gin.Context) {
	role, _ := c.Get("Roles")
	email, _ := c.Get("Email")
	var programListRes response.ProgramListResponse
	user, err := p.user.SearchEmail(email.(string))
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	for _, v := range role.([]string) {
		switch v {
		case "admin":
			programListRes.Admin, err = p.uc.GetByRole(v, user.ID)
			if err != nil {
				p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		case "panelist":
			programListRes.Panelist, err = p.uc.GetByRole(v, user.ID)
			if err != nil {
				p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		case "mentor":
			programListRes.Mentor, err = p.uc.GetByRole(v, user.ID)
			if err != nil {
				p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		case "participant":
			programListRes.Participant, err = p.uc.GetByRole(v, user.ID)
			if err != nil {
				p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, programListRes, "OK")
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

func (p *ProgramController) getQuestionHandler(c *gin.Context) {
	id := c.Param("id")
	payload, err := p.uc.ListQuestions(id)
	if err != nil {
		p.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	p.NewSuccessSingleResponse(c, payload.EvaluationCategories, "OK")
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

func NewProgramController(r *gin.Engine, auth gin.IRoutes, uc usecase.ProgramUsecase, user usecase.UserUsecase) *ProgramController {
	controller := ProgramController{
		router: r,
		auth:   auth,
		uc:     uc,
		user:   user,
	}
	auth.GET("/programs", controller.listHandler)
	r.GET("/programs/questions/:id", controller.getQuestionHandler)
	r.GET("/programs/:id", controller.getHandler)
	r.POST("/programs", controller.createHandler)
	r.PUT("/programs", controller.updateHandler)
	r.DELETE("/programs/:id", controller.deleteHandler)
	return &controller
}
