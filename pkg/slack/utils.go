package slack

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

func SendMessage(webhook, username, iconUrl, channel, color, preText, footerText string, fields []slack.AttachmentField) error {
	ts := json.Number(strconv.FormatInt(time.Now().Unix(), 10))
	attachment := slack.Attachment{
		Color:  color,
		Text:   preText,
		Footer: footerText,
		Ts:     ts,
		Fields: fields,
	}

	msg := slack.WebhookMessage{
		Username:    username,
		IconURL:     iconUrl,
		Attachments: []slack.Attachment{attachment},
		Channel:     channel,
	}

	if err := slack.PostWebhook(webhook, &msg); err != nil {
		return err
	}

	return nil
}
