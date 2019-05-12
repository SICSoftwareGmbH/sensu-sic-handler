// Copyright © 2019 SIC! Software GmbH

package recipient

// OutputType type of recipient
type OutputType int

// Output types
const (
	OutputTypeNone  OutputType = 0
	OutputTypeMail  OutputType = 1
	OutputTypeXMPP  OutputType = 2
	OutputTypeSlack OutputType = 3
)

// Recipient recipient for notifications
type Recipient struct {
	Type OutputType
	ID   string
	Args map[string]string
}
