package controller

import (
	"encoding/json"
	"net/http"

	"github.com/alwinihza/talent-connect-be/delivery/api"
	"github.com/alwinihza/talent-connect-be/model"
	"github.com/alwinihza/talent-connect-be/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	router *gin.Engine
	uc     usecase.UserUsecase
	api.BaseApi
}

func (u *UserController) listHandler(c *gin.Context) {
	searchBy := c.Query("role")
	var farmers []model.User
	var err error

	if searchBy == "" {
		farmers, err = u.uc.FindAll()
	} else {
		farmers, err = u.uc.SearchByRole(searchBy)
	}

	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, farmers, "OK")
}

func (u *UserController) searchMenteeHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	mentor_id := c.Param("mentorId")
	program_id := c.Param("id")

	payload, err := u.uc.SearchAvailableMenteeForMentor(mentor_id, program_id, name)
	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, payload, "OK")
}

func (u *UserController) searchMenteeForJudgesHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	panelist_id := c.Param("panelistId")
	program_id := c.Param("id")

	payload, err := u.uc.SearchAvailableMenteeForJudges(panelist_id, program_id, name)
	if err != nil {
		u.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	u.NewSuccessSingleResponse(c, payload, "OK")
}

func (r *UserController) createHandler(c *gin.Context) {
	user := c.PostForm("user")
	role := c.PostForm("role")
	var payload model.User
	var roleUser []string

	if err := json.Unmarshal([]byte(user), &payload); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := json.Unmarshal([]byte(role), &roleUser); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, "Role not valid")
		return
	}

	if err := r.uc.UpdateRole(&payload, roleUser); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := r.uc.SaveData(&payload); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r.NewSuccessSingleResponse(c, payload, "OK")
}

func (r *UserController) updateHandler(c *gin.Context) {
	user := c.PostForm("user")
	role := c.PostForm("role")
	var payload model.User
	var roleUser []string

	if err := json.Unmarshal([]byte(user), &payload); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, "User not valid")
		return
	}
	if err := json.Unmarshal([]byte(role), &roleUser); err != nil {
		r.NewFailedResponse(c, http.StatusBadRequest, "Role not valid")
		return
	}

	if len(roleUser) > 0 {
		if err := r.uc.UpdateRole(&payload, roleUser); err != nil {
			r.NewFailedResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	if err := r.uc.UpdateData(&payload); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r.NewSuccessSingleResponse(c, payload, "OK")
}

func (r *UserController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := r.uc.DeleteData(id); err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusNoContent, "")
}

func (r *UserController) getHandler(c *gin.Context) {
	id := c.Param("id")
	payload, err := r.uc.FindById(id)
	if err != nil {
		r.NewFailedResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	r.NewSuccessSingleResponse(c, payload, "OK")
}

func NewUserController(r *gin.Engine, uc usecase.UserUsecase) *UserController {
	controller := UserController{
		router: r,
		uc:     uc,
	}
	r.GET("/users", controller.listHandler)
	r.GET("/users/:id", controller.getHandler)
	r.PUT("/users", controller.updateHandler)
	r.POST("/users", controller.createHandler)
	r.GET("/users/mentor/:id/:mentorId", controller.searchMenteeHandler)
	r.GET("/users/panelist/:id/:panelistId", controller.searchMenteeForJudgesHandler)
	// r.DELETE("/farmers/:id", controller.deleteHandler)
	return &controller
}
