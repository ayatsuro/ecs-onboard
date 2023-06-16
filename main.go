package main

import (
	"ecs-onboard/handler"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
)

func main() {
	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		slog.Info(ctx.Request.Method, ctx.Request.URL)
		ctx.Next()
	})
	r.POST("/onboard", handler.OnboardNs)
	r.POST("/migrate", handler.MigrateNs)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
