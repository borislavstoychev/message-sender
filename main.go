package main

import (
	"io"
	"os"
	"sender/controllers"
	"sender/helpers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	helpers.InitLogger("message_sender")
	err := godotenv.Load(".env")
	if err != nil {
		helpers.GetLogger().WithFields(logrus.Fields{
			"Erorr": err.Error(),
		}).Error("Error loading .env")
		os.Exit(1)
	}

}

func main() {
	if os.Getenv("APP_MODE") != "DEBUG" {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = io.Discard
	}

	r := gin.New()
	r.Use(helpers.CustomRecoveryWithWriter())
	r.Use(helpers.GinFormatMiddleware())
	r.LoadHTMLGlob("views/*")
	r.Static("/assets", "./public/assets")
	r.GET("/", controllers.Home)
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
