package engine

import (
	"net/http"
	"time"

	util "github.com/bharath0292/quantdrey/pkg/utils"
	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RestEngineHandlers struct{}

type RestEngine struct {
	engine *gin.Engine
}

func NewRestEngine() *RestEngine {
	r := gin.New()

	return &RestEngine{
		engine: r,
	}
}

func (re *RestEngine) Setup(handlers RestEngineHandlers, graphqlEngine *GrapQLEngine) {
	re.engine.Use(gin.Recovery())
	re.engine.Use(ginzerolog.Logger("quantdrey"))
	re.engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With", "X-Api-Key"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	re.engine.OPTIONS("/*cors", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
	re.engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "App is healthy",
		})
	})
	re.engine.GET("/playground",
		func(ctx *gin.Context) {
			graphqlEngine.PlaygroundHandler(ctx.Writer, ctx.Request)
		})

	protectedRoute := re.engine.Group("/server")
	// protectedRoute.Use(middleware.AuthMiddleware())

	protectedRoute.POST("/query", func(ctx *gin.Context) {
		req := util.GinToHttpContext(ctx)
		graphqlEngine.ServeHTTP(ctx.Writer, req)
	})
}

func (re *RestEngine) Run(addr string) error {
	return re.engine.Run(addr)
}

func (re *RestEngine) RunTLS(addr, certPath, keyPath string) error {
	return re.engine.RunTLS(addr, certPath, keyPath)
}
