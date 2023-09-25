package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kkmhh/practice-go-gin/src/controller"
	"github.com/kkmhh/practice-go-gin/src/middlewares"
	"io"
	"net/http"
	"os"

	"github.com/kkmhh/practice-go-gin/src/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(
		gin.Recovery(),
		middlewares.Logger(),
		middlewares.BasicAuth(),
		//gindump.Dump(),
	)

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Valid."})
		}
	})

	server.Run(":8080")
}
