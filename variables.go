package main

var (
	mailgun_domain             string
	mailgun_api_key            string
	mailgun_ignore_event_types []string
	mailgun_bounce_alerts      []string

	cron_check_period string

	slack_username          string
	slack_channel_general   string
	slack_channel_emergency string
	slack_webhook           string
	slack_icon_url          string
	slack_footer_text       string
)
