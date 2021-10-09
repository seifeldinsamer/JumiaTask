package handler

import (
	"backend/api/model"
	"backend/usecase/users"
	"github.com/gin-gonic/gin"
)

type GinHandler struct {
	UserUseCase users.UseCase
}

func NewGinHandler(useCase users.UseCase) (r *gin.Engine) {
	h := &GinHandler{
		useCase,
	}
	r = gin.Default()
	r.GET("/users/:country_name", h.countryUsers)
	r.GET("/users/:country_name/count", h.totalCountryUsers)
	return r
}

func (h *GinHandler) countryUsers(c *gin.Context) {
	users := h.UserUseCase.CountryUsers(c.Param("country_name"))
	res := model.CountryUsersOutput{Data: users}
	c.JSON(200, res)
}

func (h *GinHandler) totalCountryUsers(c *gin.Context) {
	total := h.UserUseCase.TotalCountryUsers(c.Param("country_name"))
	res := model.TotalCountryUsersOutput{Total: total}
	c.JSON(200, res)
}
