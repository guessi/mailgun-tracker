package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

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

	slack_username = v.GetString("slack.username")
	if len(slack_username) <= 0 {
		log.Fatalf("missing config: slack_username")
	}

	slack_channel = v.GetString("slack.channel")
	if len(slack_channel) <= 0 {
		log.Fatalf("missing config: slack_channel")
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

func slackSendMessage(color, msg string) {
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
		Channel:     slack_channel,
	}

	if err := slack.PostWebhook(slack_webhook, &whmsg); err != nil {
		log.Println(err)
	}
}
