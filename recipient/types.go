// Copyright © 2019 SIC! Software GmbH

package recipient

// HandlerType type of recipient
type HandlerType int

// Handler types
const (
	HandlerTypeNone  HandlerType = 0
	HandlerTypeMail  HandlerType = 1
	HandlerTypeXMPP  HandlerType = 2
	HandlerTypeSlack HandlerType = 3
)

// Recipient recipient for notifications
type Recipient struct {
	Type HandlerType
	ID   string
	Args map[string]string
}
