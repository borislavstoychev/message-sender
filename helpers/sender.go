package helpers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

var client *twilio.RestClient

func initTwilioClient() {
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
}

// Send message to given phone number
func MessageSender(c *gin.Context) {
	if client == nil {
		initTwilioClient()
	}
	var data RequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	SendMessage(data.Message, data.ToPhone)

	// Send a response
	c.JSON(http.StatusOK, gin.H{"message": "Received and processed your message"})
}

func CreateVerify(c *gin.Context) {
	if client == nil {
		initTwilioClient()
	}
	var data RequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := &verify.CreateVerificationParams{}
	params.SetTo(data.ToPhone)
	// params.SetCustomFriendlyName(data.Message)
	// params.SetCustomMessage("This is your verification code. Do not share it.")
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(os.Getenv("VERIFY_SERVICE_SID"), params)
	if err != nil {
		GetLogger().WithFields(logrus.Fields{
			"Error": err.Error(),
		}).Error("Error creating vrify!")
	} else {
		GetLogger().WithFields(logrus.Fields{
			"Sid":    resp.Sid,
			"Status": resp.Status,
			"Valid":  resp.Valid,
		}).Info("Verification code have been send.")
	}
}

func SendMessage(msg string, toPhone string) {

	params := openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(os.Getenv("FROM_PHONE"))
	params.SetBody(msg)

	response, err := client.Api.CreateMessage(&params)
	if err != nil {
		GetLogger().WithFields(logrus.Fields{
			"Error": err.Error(),
		}).Error("Error creating and sending message!")
		return
	}
	GetLogger().Infof("Message SID: %s\n", *response.Sid)
}

func Verify(c *gin.Context) {
	if client == nil {
		initTwilioClient()
	}
	var data RequestData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(data.ToPhone)
	params.SetCode(data.Message)

	resp, err := client.VerifyV2.CreateVerificationCheck(os.Getenv("VERIFY_SERVICE_SID"), params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"response": resp})
	}
}
