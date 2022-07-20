package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nguyenthanhworkspace/golang-starter/internal/entity"
	"github.com/nguyenthanhworkspace/golang-starter/internal/usecase"
	"github.com/nguyenthanhworkspace/golang-starter/pkg/logger"
	"net/http"
)

type userRoutes struct {
	t usecase.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, u usecase.User, l logger.Interface) {
	r := &userRoutes{u, l}

	h := handler.Group("/users")
	{
		h.GET("/index", r.usersIndex)
		h.GET("/:id", r.usersIndex)
	}
}

type userIndexResponse struct {
	UserResource []entity.User `json:"userIndex"`
}

// @Summary     Show user
// @Description Show all users
// @ID          usersIndex
// @Tags  	    users
// @Accept      json
// @Produce     json
// @Success     200 {object} userIndexResponse
// @Failure     500 {object} response
// @Router      /users/index [get]
func (r *userRoutes) usersIndex(c *gin.Context) {
	c.JSON(http.StatusOK, userIndexResponse{})
}

func (r *userRoutes) userDetail(c *gin.Context) {
	c.JSON(http.StatusOK, userIndexResponse{})
}
