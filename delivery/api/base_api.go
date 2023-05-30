package api

import (
	"github.com/alwinihza/talent-connect-be/delivery/api/response"
	"github.com/gin-gonic/gin"
)

type BaseApi struct{}

func (b *BaseApi) ParseRequestBody(c *gin.Context, payload interface{}) error {
	if err := c.ShouldBindJSON(payload); err != nil {
		return err
	}
	return nil
}

func (b *BaseApi) NewSuccessSingleResponse(c *gin.Context, data interface{}, desc string) {
	response.SendSingleResponse(c, data, desc)
}

func (b *BaseApi) NewFailedResponse(c *gin.Context, code int, desc string) {
	response.SendErrorResponse(c, code, desc)
}

// func (b *BaseApi) NewSuccessPagedResponse(c *gin.Context, data interface{}, desc string, paging dto.Paging) {
// 	response.SendPagedResponse(c, data, desc, paging)
// }
