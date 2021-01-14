package config

type Config struct {
	MailgunConfig Mailgun `mapstructure:"mailgun,omitempty"`
	SlackConfig   Slack   `mapstructure:"slack,omitempty"`
	CronConfig    Cron    `mapstructure:"cron,omitempty"`
}

type Mailgun struct {
	Domain                  string   `yaml:"domain,omitempty"`
	APIKey                  string   `yaml:"apiKey,omitempty"`
	IgnoreEventTypes        []string `yaml:"ignoreEventTypes,omitempty"`
	IgnoreRecipientKeywords []string `yaml:"ignoreRecipientKeywords,omitempty"`
	IgnoreSubjectKeywords   []string `yaml:"ignoreSubjectKeywords,omitempty"`
	BounceAlerts            []string `yaml:"bounceAlerts,omitempty"`
}

type Slack struct {
	Username         string `yaml:"username,omitempty"`
	ChannelGeneral   string `yaml:"channelGeneral,omitempty"`
	ChannelEmergency string `yaml:"channelEmergency,omitempty"`
	Webhook          string `yaml:"webhook,omitempty"`
	IconUrl          string `yaml:"iconUrl,omitempty"`
	FooterText       string `yaml:"footerText,omitempty"`
}

type Cron struct {
	CheckPeriod string `yaml:"checkPeriod,omitempty"`
}
