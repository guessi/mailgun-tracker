package mailgun

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guessi/mailgun-tracker/config"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/mailgun/mailgun-go/v4/events"
)

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func CheckBounce(m config.Mailgun, s config.Slack) {
	mg := mailgun.NewMailgun(m.Domain, m.APIKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for _, bounce := range m.BounceAlerts {
		bounce, err := mg.GetBounce(ctx, bounce)
		if err != nil {
			// skip if nothing found
			continue
		}
		bounceHandler(bounce, s)
	}
}

func PermanentFailureHandler(ctx *gin.Context, m config.Mailgun, s config.Slack) {
	var payload mailgun.WebhookPayload

	if err := json.NewDecoder(ctx.Request.Body).Decode(&payload); err != nil {
		fmt.Printf("decode JSON error: %s", err)
		ctx.JSON(http.StatusNotAcceptable, "error")
		return
	}

	if err := verifyWebhookSignature(m.Domain, m.APIKey, payload.Signature); err != nil {
		ctx.JSON(http.StatusNotAcceptable, "webhook signature verification failed")
	}

	e, err := mailgun.ParseEvent(payload.EventData)
	if err != nil {
		fmt.Printf("parse event error: %s\n", err)
		return
	}

	switch event := e.(type) {
	case *events.Failed:
		eventFailedHandler(m, s, event)
	}

	// FIXME: should handle return message
	ctx.JSON(http.StatusOK, gin.H{
		"status": "posted",
	})
}
