package main

import (
	_ "ecs-onboard/docs"
	"ecs-onboard/handler"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /v1

func main() {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		slog.Info(ctx.Request.Method, ctx.Request.URL)
		ctx.Next()
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/v1")
	v1.POST("/namespace/onboard", handler.OnboardNs)
	v1.POST("/namespace/migrate", handler.MigrateNs)
	v1.GET("/test", handler.Test)

	if err := r.Run("0.0.0.0:8081"); err != nil {
		slog.Fatal(err)
	}
}