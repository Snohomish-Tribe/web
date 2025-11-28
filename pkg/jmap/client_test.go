package jmap

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// TestMain runs before all tests and loads the .env file
func TestMain(m *testing.M) {
	// Load .env file from project root (two directories up from
	// pkg/jmap)
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("Warning: .env file not found or couldn't be loaded: %v", err)
		log.Printf("Integration tests will be skipped unless environment variables are set manually")
	}

	// Run all tests
	code := m.Run()

	// Exit with the test result code
	os.Exit(code)
}

// TestNewClient tests creating a new JMAP client
// Requires FASTMAIL_TOKEN environment variable
func TestNewClient(t *testing.T) {
	token := os.Getenv("FASTMAIL_TOKEN")
	if token == "" {
		t.Skip("FASTMAIL_TOKEN not set, skipping integration test")
	}

	client, err := NewClient(
		"https://api.fastmail.com/jmap/session",
		token,
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	if client == nil {
		t.Fatal("Client is nil")
	}

	if !client.initialized {
		t.Error("Client not properly initialized")
	}

	if client.accountID == "" {
		t.Error("Account ID not set")
	}

	if client.draftsID == "" {
		t.Error("Drafts mailbox ID not set")
	}

	if client.sentID == "" {
		t.Error("Sent mailbox ID not set")
	}
}

// TestNewClientInvalidToken tests authentication with an invalid token
func TestNewClientInvalidToken(t *testing.T) {
	_, err := NewClient(
		"https://api.fastmail.com/jmap/session",
		"invalid-token",
	)
	if err == nil {
		t.Error("Expected error with invalid token, got nil")
	}
}

// TestSendEmail tests sending a plain text email
// Requires FASTMAIL_TOKEN and TEST_SENDER_EMAIL environment
// variables
func TestSendEmail(t *testing.T) {
	token := os.Getenv("FASTMAIL_TOKEN")
	senderEmail := os.Getenv("TEST_SENDER_EMAIL")

	if token == "" || senderEmail == "" {
		t.Skip("FASTMAIL_TOKEN and/or TEST_SENDER_EMAIL not set, skipping integration test")
	}

	client, err := NewClient(
		"https://api.fastmail.com/jmap/session",
		token,
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	from := NewAddress("Test Sender", senderEmail)
	to := []Address{
		NewAddress("Test Recipient", senderEmail), // Send to self
	}

	err = client.SendEmail(
		from,
		to,
		"Test Email from JMAP Client",
		"This is a test email sent from the JMAP client test suite.",
		false,
	)
	if err != nil {
		t.Fatalf("Failed to send email: %v", err)
	}
}

// TestSendEmailMultipleRecipients tests sending an email to multiple
// recipients
func TestSendEmailMultipleRecipients(t *testing.T) {
	token := os.Getenv("FASTMAIL_TOKEN")
	senderEmail := os.Getenv("TEST_SENDER_EMAIL")
	recipientEmail := os.Getenv("TEST_RECIPIENT_EMAIL")

	if token == "" || senderEmail == "" {
		t.Skip("FASTMAIL_TOKEN and/or TEST_SENDER_EMAIL not set, skipping integration test")
	}

	// Use sender email as recipient if TEST_RECIPIENT_EMAIL not set
	if recipientEmail == "" {
		recipientEmail = senderEmail
	}

	client, err := NewClient(
		"https://api.fastmail.com/jmap/session",
		token,
	)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	from := NewAddress("Test Sender", senderEmail)
	to := []Address{
		NewAddress("First Recipient", senderEmail),
		NewAddress("Second Recipient", recipientEmail),
	}

	err = client.SendEmail(
		from,
		to,
		"Test Email with Multiple Recipients",
		"This is a test email sent to multiple recipients.",
		false,
	)
	if err != nil {
		t.Fatalf("Failed to send email to multiple recipients: %v", err)
	}
}

// TestNewAddress tests the NewAddress helper function
func TestNewAddress(t *testing.T) {
	addr := NewAddress("John Doe", "john@example.com")

	if addr.Name != "John Doe" {
		t.Errorf("Expected name 'John Doe', got '%s'", addr.Name)
	}

	if addr.Email != "john@example.com" {
		t.Errorf("Expected email 'john@example.com', got '%s'",
			addr.Email)
	}
}

// TestAddressConversion tests the helper functions for converting
// addresses
func TestAddressConversion(t *testing.T) {
	addresses := []Address{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
	}

	emailAddrs := addressesToEmailAddresses(addresses)
	if len(emailAddrs) != 2 {
		t.Fatalf("Expected 2 email addresses, got %d", len(emailAddrs))
	}

	if emailAddrs[0].Name != "Alice" || emailAddrs[0].Email != "alice@example.com" {
		t.Errorf("First address mismatch: %+v", emailAddrs[0])
	}

	if emailAddrs[1].Name != "Bob" || emailAddrs[1].Email != "bob@example.com" {
		t.Errorf("Second address mismatch: %+v", emailAddrs[1])
	}

	submissionAddrs := addressesToSubmissionAddresses(addresses)
	if len(submissionAddrs) != 2 {
		t.Fatalf("Expected 2 submission addresses, got %d",
			len(submissionAddrs))
	}

	if submissionAddrs[0].Email != "alice@example.com" {
		t.Errorf("First submission address mismatch: %s",
			submissionAddrs[0].Email)
	}

	if submissionAddrs[1].Email != "bob@example.com" {
		t.Errorf("Second submission address mismatch: %s",
			submissionAddrs[1].Email)
	}
}
