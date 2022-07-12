package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
//func NewRouter(r *gin.Engine, l logger.Interface) {
func NewRouter(r *gin.Engine) {
	// Options
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Swagger
	//swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	//handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	r.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	//r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := r.Group("api/v1")
	{
		newBookRoutes(h)
	}
}
