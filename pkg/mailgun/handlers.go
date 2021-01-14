package mailgun

import (
	"fmt"
	"log"
	"strings"

	"github.com/guessi/mailgun-tracker/config"
	slackmsg "github.com/guessi/mailgun-tracker/pkg/slack"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/mailgun/mailgun-go/v4/events"
	"github.com/slack-go/slack"
)

func bounceHandler(b mailgun.Bounce, s config.Slack) {
	preText := "<!channel> :bangbang: :bangbang: :bangbang:"
	attachmentFields := []slack.AttachmentField{
		slack.AttachmentField{
			Title: "Alert Infomation",
			Value: fmt.Sprintf("Receiver `%s` was blocked by mailgun", b.Address),
			Short: false,
		},
		slack.AttachmentField{
			Title: "Detail Information",
			Value: b.Error,
			Short: false,
		},
		slack.AttachmentField{
			Title: "Error Code",
			Value: b.Code,
			Short: true,
		},
		slack.AttachmentField{
			Title: "Created At",
			Value: b.CreatedAt.String(),
			Short: true,
		},
	}

	if err := slackmsg.SendMessage(
		s.Webhook, s.Username, s.IconUrl, s.ChannelEmergency, "danger",
		preText, s.FooterText, attachmentFields,
	); err != nil {
		log.Printf("Failed to send slack message, err: %v", err)
	}
}

func skipFailureOrNot(m config.Mailgun, e *events.Failed) bool {
	for _, t := range m.IgnoreEventTypes {
		if e.Reason == t {
			log.Printf("Skip event with event type: %s", t)
			return true
		}
	}

	for _, r := range m.IgnoreRecipientKeywords {
		if strings.Contains(e.Message.Headers.To, r) {
			log.Printf("Skip recipient %s contains %s", e.Message.Headers.To, r)
			return true
		}
	}

	for _, s := range m.IgnoreSubjectKeywords {
		if strings.Contains(e.Message.Headers.Subject, s) {
			log.Printf("Skip subject %s contains %s", e.Message.Headers.Subject, s)
			return true
		}
	}

	return false
}

func eventFailedHandler(m config.Mailgun, s config.Slack, event *events.Failed) {
	if skipFailureOrNot(m, event) {
		return
	}

	preText := "<!here> :bangbang: :bangbang: :bangbang:"
	attachmentFields := []slack.AttachmentField{
		slack.AttachmentField{
			Title: "Subject",
			Value: event.Message.Headers.Subject,
			Short: false,
		},
		slack.AttachmentField{
			Title: "From",
			Value: event.Message.Headers.From,
			Short: true,
		},
		slack.AttachmentField{
			Title: "To",
			Value: event.Message.Headers.To,
			Short: true,
		},
		slack.AttachmentField{
			Title: "Message ID",
			Value: event.Message.Headers.MessageID,
			Short: false,
		},
		slack.AttachmentField{
			Title: "Serverity",
			Value: event.Severity,
			Short: true,
		},
		slack.AttachmentField{
			Title: "Status Code",
			Value: fmt.Sprintf("%d", event.DeliveryStatus.Code),
			Short: true,
		},
		slack.AttachmentField{
			Title: "Reason",
			Value: event.Reason,
			Short: false,
		},
	}

	if err := slackmsg.SendMessage(
		s.Webhook, s.Username, s.IconUrl, s.ChannelGeneral, "warning",
		preText, s.FooterText, attachmentFields,
	); err != nil {
		log.Printf("Failed to send slack message, err: %v", err)
	}
}
