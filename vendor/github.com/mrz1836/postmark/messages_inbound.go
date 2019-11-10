package postmark

import (
	"fmt"
	"net/url"
	"time"
)

// InboundMessage - a message received from the Postmark server
type InboundMessage struct {
	// From - The sender email address.
	From string
	// FromName - The sender name.
	FromName string
	// FromFull - Sender email address and name.
	FromFull Recipient
	// To - Inbound address the message was sent to.
	To string
	// ToFull - Slice of all TO recipients
	ToFull []Recipient
	// CcFull - Slice of all CC recipients
	CcFull []Recipient
	// Cc - Cc recipient email address.
	Cc string
	// ReplyTo - Reply to override email address.
	ReplyTo string
	// OriginalRecipient - Receiver (RCPT TO) address this webhook is for.
	OriginalRecipient string
	// Subject - Email subject
	Subject string
	// Date - Timestamp
	Date string
	// MailboxHash - Custom hash that the email was sent to.
	MailboxHash string
	// TextBody - Plain text email message.
	TextBody string
	// HTMLBody - HTML email message.
	HTMLBody string `json:"HtmlBody"`
	// Tag - Tag name
	Tag string
	// Headers - List of objects that each represent a header name and value.
	Headers []Header
	// Attachments - List of objects that each represent an attachment.
	Attachments []Attachment
	// MessageID - Unique ID of the message.
	MessageID string
	// BlockedReason - Reason message was blocked.
	BlockedReason string
	// Status - Status of message in your Postmark activity.
	Status string
}

// Time returns a parsed time.Time struct
// Inbound messages return as RFC1123Z strangely
func (x InboundMessage) Time() (time.Time, error) {
	return time.Parse(time.RFC1123Z, x.Date)
}

///////////////////////////////////////
///////////////////////////////////////

// GetInboundMessage fetches a specific inbound message via serverID
func (client *Client) GetInboundMessage(messageID string) (InboundMessage, error) {
	res := InboundMessage{}
	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("messages/inbound/%s/details", messageID),
		TokenType: serverToken,
	}, &res)
	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

type inboundMessagesResponse struct {
	TotalCount int64
	Messages   []InboundMessage
}

// GetInboundMessages fetches a list of inbound message on the server
// It returns a InboundMessage slice, the total message count, and any error that occurred
// http://developer.postmarkapp.com/developer-api-messages.html#inbound-message-search
func (client *Client) GetInboundMessages(count int64, offset int64, options map[string]interface{}) ([]InboundMessage, int64, error) {
	res := inboundMessagesResponse{}

	values := &url.Values{}
	values.Add("count", fmt.Sprintf("%d", count))
	values.Add("offset", fmt.Sprintf("%d", offset))

	for k, v := range options {
		values.Add(k, fmt.Sprintf("%v", v))
	}

	err := client.doRequest(parameters{
		Method:    "GET",
		Path:      fmt.Sprintf("messages/inbound?%s", values.Encode()),
		TokenType: serverToken,
	}, &res)

	return res.Messages, res.TotalCount, err
}

///////////////////////////////////////
///////////////////////////////////////

// BypassInboundMessage - Bypass rules for a blocked inbound message
func (client *Client) BypassInboundMessage(messageID string) error {
	res := APIError{}
	err := client.doRequest(parameters{
		Method:    "PUT",
		Path:      fmt.Sprintf("messages/inbound/%s/bypass", messageID),
		TokenType: serverToken,
	}, &res)

	if res.ErrorCode != 0 {
		return res
	}

	return err
}

///////////////////////////////////////
///////////////////////////////////////

// RetryInboundMessage - Retry a failed inbound message for processing
func (client *Client) RetryInboundMessage(messageID string) error {
	res := APIError{}
	err := client.doRequest(parameters{
		Method:    "PUT",
		Path:      fmt.Sprintf("messages/inbound/%s/retry", messageID),
		TokenType: serverToken,
	}, &res)

	if res.ErrorCode != 0 {
		return res
	}

	return err
}
