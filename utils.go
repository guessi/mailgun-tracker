package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	mailgun "github.com/mailgun/mailgun-go/v4"
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

func loadConfigure(v *viper.Viper) {
	// set config path
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	// check required configs
	mailgun_domain = v.GetString("mailgun.domain")
	if len(mailgun_domain) <= 0 {
		log.Fatalf("missing config: mailgun_domain")
	}

	mailgun_api_key = v.GetString("mailgun.api_key")
	if len(mailgun_api_key) <= 0 {
		log.Fatalf("missing config: mailgun_api_key")
	}

	mailgun_ignore_event_types = v.GetStringSlice("mailgun.ignore_event_types")
	if len(mailgun_ignore_event_types) <= 0 {
		log.Printf("missing config: mailgun_ignore_event_types")
	}

	mailgun_bounce_alerts = v.GetStringSlice("mailgun.bounce_alerts")
	if len(mailgun_bounce_alerts) <= 0 {
		log.Printf("missing config: mailgun_bounce_alerts")
	}

	cron_check_period = v.GetString("cron.check_period")
	if len(cron_check_period) <= 0 {
		log.Fatalf("missing config: cron_check_period")
	}

	slack_username = v.GetString("slack.username")
	if len(slack_username) <= 0 {
		log.Fatalf("missing config: slack_username")
	}

	slack_channel_general = v.GetString("slack.channel_general")
	if len(slack_channel_general) <= 0 {
		log.Fatalf("missing config: slack_channel_general")
	}

	slack_channel_emergency = v.GetString("slack.channel_emergency")
	if len(slack_channel_emergency) <= 0 {
		log.Fatalf("missing config: slack_channel_emergency")
	}

	slack_webhook = v.GetString("slack.webhook")
	if len(slack_webhook) <= 0 {
		log.Fatalf("missing config: slack_webhook")
	}

	slack_icon_url = v.GetString("slack.icon_url")
	if len(slack_icon_url) <= 0 {
		log.Fatalf("missing config: slack_icon_url")
	}

	slack_footer_text = v.GetString("slack.footer_text")
	if len(slack_footer_text) <= 0 {
		log.Fatalf("missing config: slack_footer_text")
	}

	// set default values
	v.SetDefault("port", "8080")
}

func slackSendMessage(channel, color, msg string) {
	attachment := slack.Attachment{
		Color:    color,
		Fallback: msg,
		Text:     "<!here> " + msg,
		Footer:   slack_footer_text,
		Ts:       json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	}

	whmsg := slack.WebhookMessage{
		Username:    slack_username,
		IconURL:     slack_icon_url,
		Attachments: []slack.Attachment{attachment},
		Channel:     channel,
	}

	if err := slack.PostWebhook(slack_webhook, &whmsg); err != nil {
		log.Println(err)
	}
}

func checkBounce() {
	mg := mailgun.NewMailgun(mailgun_domain, mailgun_api_key)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	for _, bounce := range mailgun_bounce_alerts {
		s, err := mg.GetBounce(ctx, bounce)
		if err != nil {
			continue // skip if nothing found
		}

		msg := fmt.Sprintf(":scream::scream::scream:\n*Message:* %s was blocked by mailgun\n*Error Code*: %s\n*Error*: %s", s.Address, s.Code, s.Error)
		slackSendMessage(slack_channel_emergency, "danger", msg)
	}
}
