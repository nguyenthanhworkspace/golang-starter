// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/nguyenthanhworkspace/golang-starter/internal/usecase"
	"github.com/nguyenthanhworkspace/golang-starter/pkg/httpresponse"
	"github.com/nguyenthanhworkspace/golang-starter/pkg/logger"

	// Swagger docs.
	_ "github.com/nguyenthanhworkspace/golang-starter/docs"
)

// UseCaseList =.
type UseCaseList struct {
	translation usecase.Translation
	user        usecase.User
}

func NewUseCaseList(t usecase.Translation, u usecase.User) *UseCaseList {
	return &UseCaseList{
		translation: t,
		user:        u,
	}
}

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, uc *UseCaseList) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) {
		res := httpresponse.BuildResponse(true, "Ping", nil)
		c.JSON(http.StatusOK, res)
	})

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newTranslationRoutes(h, uc.translation, l)
		newUserRoutes(h, uc.user, l)
	}
}
