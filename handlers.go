package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	mailgun "github.com/mailgun/mailgun-go/v4"
	"github.com/mailgun/mailgun-go/v4/events"
)

func verifyWebhookSignature(s mailgun.Signature) error {
	mg := mailgun.NewMailgun(mailgun_domain, mailgun_api_key)

	verified, err := mg.VerifyWebhookSignature(s)
	if err != nil {
		return err
	}

	if !verified {
		return errors.New("webhook signature verification failed")
	}

	return nil
}

func logFailedAndSendNotification(event *events.Failed) {
	msg := fmt.Sprintf("*StatusCode:* %d, *Severity:* %s, *Reason:* %s\n*Subject:* %s\n*From:* %s\n*To:* %s\n*MessageId:* %s\n",
		event.DeliveryStatus.Code,
		event.Severity,
		event.Reason,
		event.Message.Headers.Subject,
		event.Message.Headers.From,
		event.Message.Headers.To,
		event.Message.Headers.MessageID,
	)

	slackSendMessage("warning", msg)
}

func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func permanentFailureHandler(c *gin.Context) {
	var payload mailgun.WebhookPayload

	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		fmt.Printf("decode JSON error: %s", err)
		c.JSON(http.StatusNotAcceptable, "error")
		return
	}

	if err := verifyWebhookSignature(payload.Signature); err != nil {
		c.JSON(http.StatusNotAcceptable, "webhook signature verification failed")
	}

	e, err := mailgun.ParseEvent(payload.EventData)
	if err != nil {
		fmt.Printf("parse event error: %s\n", err)
		return
	}

	switch event := e.(type) {
	case *events.Failed:
		logFailedAndSendNotification(event)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "posted",
	})
}
