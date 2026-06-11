package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/cwinters8/gomap"
	"github.com/cwinters8/gomap/objects/emails"
	"github.com/cwinters8/gomap/requests"
	"github.com/google/uuid"
)

func sendContactEmail(mail *gomap.Client, from, to, replyTo *emails.Address, subject, body string) error {
	email, err := emails.NewEmail(
		[]string{mail.Drafts.ID},
		[]*emails.Address{from},
		[]*emails.Address{to},
		subject,
		body,
		emails.TextPlain,
	)
	if err != nil {
		return fmt.Errorf("failed to instantiate contact email: %w", err)
	}

	call, err := newEmailSetWithReplyToCall(mail.Session.PrimaryAccounts.Mail, email, replyTo)
	if err != nil {
		return err
	}

	responses, err := requests.Request(mail.Client, []*requests.Call{call}, false)
	if err != nil {
		return fmt.Errorf("email set request failure: %w", err)
	}
	if len(responses) < 1 {
		return fmt.Errorf("email set returned no responses")
	}

	if err := setCreatedEmailID(email, responses[0].Body); err != nil {
		return err
	}
	if _, err := email.Submit(mail.Client, mail.Drafts.ID, mail.Sent.ID); err != nil {
		return fmt.Errorf("failed to submit contact email: %w", err)
	}
	return nil
}

func newEmailSetWithReplyToCall(accountID string, email *emails.Email, replyTo *emails.Address) (*requests.Call, error) {
	rawEmail, err := json.Marshal(email)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal contact email: %w", err)
	}

	var emailPayload map[string]any
	if err := json.Unmarshal(rawEmail, &emailPayload); err != nil {
		return nil, fmt.Errorf("failed to build contact email payload: %w", err)
	}
	emailPayload["replyTo"] = []*emails.Address{replyTo}

	callID, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate email set call id: %w", err)
	}
	return &requests.Call{
		ID:        callID,
		AccountID: accountID,
		Method:    "Email/set",
		Arguments: map[string]any{
			"create": map[string]any{
				email.RequestID.String(): emailPayload,
			},
		},
	}, nil
}

func setCreatedEmailID(email *emails.Email, body map[string]any) error {
	rawBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal email set response: %w", err)
	}

	var response struct {
		Created map[string]struct {
			ID string `json:"id"`
		} `json:"created"`
		NotCreated map[string]struct {
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"notCreated"`
	}
	if err := json.Unmarshal(rawBody, &response); err != nil {
		return fmt.Errorf("failed to unmarshal email set response: %w", err)
	}

	requestID := email.RequestID.String()
	if failure, ok := response.NotCreated[requestID]; ok {
		return fmt.Errorf("email creation failed with error type %q and description %q", failure.Type, failure.Description)
	}
	created, ok := response.Created[requestID]
	if !ok || created.ID == "" {
		return fmt.Errorf("created email id not found for request %s", requestID)
	}
	email.ID = created.ID
	return nil
}
