package fcm

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/mcbattirola/notify-me/server/models"
)

const fcmEndpoint = "https://fcm.googleapis.com/fcm/send"

func SendPushNotification(notification models.Notification, deviceID string) {
	fcmAuthToken := os.Getenv("FCM_KEY")

	jsonStream := `
	{
		"to": "` + deviceID + `",
		"notification": {
			"body": "` + notification.Body + `",
			"title": "` + notification.Title + `"
		}
	}
	`

	requestBody := ioutil.NopCloser(strings.NewReader(jsonStream))

	req, err := http.NewRequest("POST", fcmEndpoint, requestBody)

	authorizationToken := "key = " + fcmAuthToken
	req.Header.Add("Authorization", authorizationToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	_, err = client.Do(req)
	if err != nil {
		println("Error on response.\n[ERRO] -", err)
	}

}
