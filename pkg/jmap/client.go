package jmap

import (
	"fmt"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
	"git.sr.ht/~rockorager/go-jmap/mail/email"
	"git.sr.ht/~rockorager/go-jmap/mail/emailsubmission"
	"git.sr.ht/~rockorager/go-jmap/mail/mailbox"
)

// Client wraps the JMAP client with helper methods for sending emails
type Client struct {
	client      *jmap.Client
	accountID   jmap.ID
	draftsID    jmap.ID
	sentID      jmap.ID
	initialized bool
}

// Address represents an email address with an optional name
type Address struct {
	Name  string
	Email string
}

// NewAddress creates a new Address
func NewAddress(name, email string) Address {
	return Address{Name: name, Email: email}
}

// NewClient creates a new JMAP client for Fastmail
func NewClient(sessionURL, accessToken string) (*Client, error) {
	client := &jmap.Client{
		SessionEndpoint: sessionURL,
	}
	client.WithAccessToken(accessToken)

	// Authenticate to get the session
	if err := client.Authenticate(); err != nil {
		return nil, fmt.Errorf("failed to authenticate: %w", err)
	}

	// Get the primary mail account ID
	accountID, ok := client.Session.PrimaryAccounts[mail.URI]
	if !ok {
		return nil, fmt.Errorf("no primary mail account found")
	}

	c := &Client{
		client:    client,
		accountID: accountID,
	}

	// Get mailbox IDs for Drafts and Sent
	if err := c.initializeMailboxes(); err != nil {
		return nil, fmt.Errorf("failed to initialize mailboxes: %w", err)
	}

	return c, nil
}

// initializeMailboxes fetches the Drafts and Sent mailbox IDs
func (c *Client) initializeMailboxes() error {
	req := &jmap.Request{}
	req.Invoke(&mailbox.Get{
		Account: c.accountID,
	})

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to get mailboxes: %w", err)
	}

	for _, inv := range resp.Responses {
		if r, ok := inv.Args.(*mailbox.GetResponse); ok {
			for _, mbox := range r.List {
				switch mbox.Role {
				case mailbox.RoleDrafts:
					c.draftsID = mbox.ID
				case mailbox.RoleSent:
					c.sentID = mbox.ID
				}
			}
		}
	}

	if c.draftsID == "" {
		return fmt.Errorf("drafts mailbox not found")
	}
	if c.sentID == "" {
		return fmt.Errorf("sent mailbox not found")
	}

	c.initialized = true
	return nil
}

// SendEmail sends an email using JMAP
// from: sender address (single address)
// to: recipient addresses (one or more)
// subject: email subject
// body: plain text email body
// isHTML: whether the body is HTML (currently only plain text is
// supported)
func (c *Client) SendEmail(from Address, to []Address, subject, body string, isHTML bool) error {
	if !c.initialized {
		return fmt.Errorf("client not properly initialized")
	}

	// Build the email object
	// Note: We need to construct the body structure properly.
	// JMAP requires both the structure (bodyStructure) and the
	// values (bodyValues)
	bodyType := "text/plain"
	if isHTML {
		bodyType = "text/html"
	}

	// Create the body part that will be referenced
	bodyPart := &email.BodyPart{
		PartID:  "1",
		Type:    bodyType,
		Charset: "UTF-8",
	}

	eml := &email.Email{
		MailboxIDs: map[jmap.ID]bool{
			c.draftsID: true,
		},
		From: []*mail.Address{
			{
				Name:  from.Name,
				Email: from.Email,
			},
		},
		To:            addressesToEmailAddresses(to),
		Subject:       subject,
		BodyStructure: bodyPart,
		BodyValues: map[string]*email.BodyValue{
			"1": {
				Value: body,
			},
		},
	}

	// Set the appropriate body reference - must match bodyStructure
	if isHTML {
		eml.HTMLBody = []*email.BodyPart{bodyPart}
	} else {
		eml.TextBody = []*email.BodyPart{bodyPart}
	}

	// Create the email draft
	req := &jmap.Request{}
	createCallID := req.Invoke(&email.Set{
		Account: c.accountID,
		Create: map[jmap.ID]*email.Email{
			"draft": eml,
		},
	})

	// Submit the email for delivery
	req.Invoke(&emailsubmission.Set{
		Account: c.accountID,
		Create: map[jmap.ID]*emailsubmission.EmailSubmission{
			"submission": {
				EmailID: jmap.ID(fmt.Sprintf("#%s", createCallID)),
				Envelope: &emailsubmission.Envelope{
					MailFrom: &emailsubmission.Address{
						Email: from.Email,
					},
					RcptTo: addressesToSubmissionAddresses(to),
				},
			},
		},
		OnSuccessDestroyEmail: []jmap.ID{jmap.ID(fmt.Sprintf("#%s", createCallID))},
	})

	// Execute the request
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	// Check for errors in the response
	for _, inv := range resp.Responses {
		switch r := inv.Args.(type) {
		case *email.SetResponse:
			if len(r.NotCreated) > 0 {
				for id, setErr := range r.NotCreated {
					errMsg := fmt.Sprintf("type: %s", setErr.Type)
					if setErr.Description != nil {
						errMsg += fmt.Sprintf(", description: %s", *setErr.Description)
					}
					if setErr.Properties != nil && len(*setErr.Properties) > 0 {
						errMsg += fmt.Sprintf(", properties: %v", *setErr.Properties)
					}
					return fmt.Errorf("failed to create email draft %s: %s",
						id, errMsg)
				}
			}
		case *emailsubmission.SetResponse:
			if len(r.NotCreated) > 0 {
				for id, setErr := range r.NotCreated {
					errMsg := fmt.Sprintf("type: %s", setErr.Type)
					if setErr.Description != nil {
						errMsg += fmt.Sprintf(", description: %s", *setErr.Description)
					}
					if setErr.Properties != nil && len(*setErr.Properties) > 0 {
						errMsg += fmt.Sprintf(", properties: %v", *setErr.Properties)
					}
					return fmt.Errorf("failed to submit email %s: %s",
						id, errMsg)
				}
			}
		case *jmap.MethodError:
			errMsg := fmt.Sprintf("type: %s", r.Type)
			if r.Description != nil {
				errMsg += fmt.Sprintf(", description: %s", *r.Description)
			}
			return fmt.Errorf("JMAP method error: %s", errMsg)
		}
	}

	return nil
}

// Helper functions to convert addresses

func addressesToEmailAddresses(addrs []Address) []*mail.Address {
	result := make([]*mail.Address, len(addrs))
	for i, addr := range addrs {
		result[i] = &mail.Address{
			Name:  addr.Name,
			Email: addr.Email,
		}
	}
	return result
}

func addressesToSubmissionAddresses(addrs []Address) []*emailsubmission.Address {
	result := make([]*emailsubmission.Address, len(addrs))
	for i, addr := range addrs {
		result[i] = &emailsubmission.Address{
			Email: addr.Email,
		}
	}
	return result
}
