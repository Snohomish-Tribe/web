package handlers

import (
	"testing"

	"github.com/cwinters8/gomap"
	"github.com/cwinters8/gomap/objects/emails"
)

func TestNewEmailSetWithReplyToCall(t *testing.T) {
	from := gomap.NewAddress("Snohomish Tribe Contact Form", "administrator@info.snohomishtribe.org")
	to := gomap.NewAddress("Snohomish Tribe Events", "kvansenus@snohomishtribe.org")
	replyTo := gomap.NewAddress("Form Submitter", "submitter@example.com")

	email, err := emails.NewEmail(
		[]string{"drafts-id"},
		[]*emails.Address{from},
		[]*emails.Address{to},
		"Contact Page Question: Events",
		"Test message",
		emails.TextPlain,
	)
	if err != nil {
		t.Fatal(err)
	}

	call, err := newEmailSetWithReplyToCall("account-id", email, replyTo)
	if err != nil {
		t.Fatal(err)
	}

	create, ok := call.Arguments["create"].(map[string]any)
	if !ok {
		t.Fatalf("create argument has type %T", call.Arguments["create"])
	}
	payload, ok := create[email.RequestID.String()].(map[string]any)
	if !ok {
		t.Fatalf("email payload has type %T", create[email.RequestID.String()])
	}

	assertPayloadAddress(t, payload, "from", from.Email)
	assertPayloadAddress(t, payload, "to", to.Email)
	assertPayloadAddress(t, payload, "replyTo", replyTo.Email)
}

func assertPayloadAddress(t *testing.T, payload map[string]any, field, expected string) {
	t.Helper()

	addresses, ok := payload[field].([]any)
	if !ok {
		if typedAddresses, typed := payload[field].([]*emails.Address); typed {
			if len(typedAddresses) == 1 && typedAddresses[0].Email == expected {
				return
			}
		}
		t.Fatalf("%s has type %T", field, payload[field])
	}
	if len(addresses) != 1 {
		t.Fatalf("%s contains %d addresses", field, len(addresses))
	}
	address, ok := addresses[0].(map[string]any)
	if !ok {
		t.Fatalf("%s address has type %T", field, addresses[0])
	}
	if address["email"] != expected {
		t.Errorf("%s email = %v, want %s", field, address["email"], expected)
	}
}
