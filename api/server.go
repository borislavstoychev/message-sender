package api

import (
	"io"
	"net/http"
	"os"
	"sender/helpers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var r *gin.Engine

func router() {
	if os.Getenv("APP_MODE") != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}

	r.Use(helpers.CustomRecoveryWithWriter())
	r.Use(helpers.GinFormatMiddleware())
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, "hello go from vercel !!!!")
	})
	r.POST("/send", helpers.MessageSender)
	r.POST("/create-verify", helpers.CreateVerify)
	r.POST("/verify", helpers.Verify)
	go func() {
		if err := r.Run(":8080"); err != nil {
			helpers.GetLogger().WithFields(logrus.Fields{
				"Error": err,
			}).Error("Error during Run")
		}
	}()
	forever := make(chan bool)
	// msg := fmt.Sprintf(os.Getenv("MSG"), "James")
	// SendMessage(msg)
	<-forever

}

func init() {
	r = gin.New()
	helpers.InitLogger("message_sender")
	err := godotenv.Load(".env")
	if err != nil {
		helpers.GetLogger().WithFields(logrus.Fields{
			"Erorr": err.Error(),
		}).Error("Error loading .env")
		os.Exit(1)
	}
	router()

}

// This script is needed for vercel
func Hendler(w http.ResponseWriter, rout *http.Request) {
	r.ServeHTTP(w, rout)
}
