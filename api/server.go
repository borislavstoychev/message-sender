package api

import (
	"io"
	"net/http"
	"os"
	"sender/helpers"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func router() {
	if os.Getenv("APP_MODE") != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}
	r = gin.New()
	r.Use(helpers.CustomRecoveryWithWriter())
	r.Use(helpers.GinFormatMiddleware())
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello go from vercel !!!!")
	})
	r.POST("/send", helpers.MessageSender)
	r.POST("/create-verify", helpers.CreateVerify)
	r.POST("/verify", helpers.Verify)
}

func init() {
	helpers.InitLogger("message_sender")
	router()
}

// This script is needed for vercel
func Handler(w http.ResponseWriter, req *http.Request) {
	r.ServeHTTP(w, req)
}
