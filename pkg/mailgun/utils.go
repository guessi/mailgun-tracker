package mailgun

import (
	"errors"

	"github.com/mailgun/mailgun-go/v4"
)

func verifyWebhookSignature(domain, apiKey string, s mailgun.Signature) error {
	mg := mailgun.NewMailgun(domain, apiKey)

	verified, err := mg.VerifyWebhookSignature(s)
	if err != nil {
		return err
	}

	if !verified {
		return errors.New("webhook signature verification failed")
	}

	return nil
}
