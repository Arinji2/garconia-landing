package logging

type Embed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Color       int     `json:"color"`
	Fields      []Field `json:"fields,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

// WebhookPayload represents the entire Discord webhook payload
type WebhookPayload struct {
	Embeds []Embed `json:"embeds"`
}
