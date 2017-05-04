package webhooks

// Webhook to model a webhook response
type Webhook struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	TargetURL string `json:"targetUrl"`
	Resource  string `json:"resource"`
	Event     string `json:"event"`
	Filter    string `json:"filter"`
	Secret    string `json:"secret"`
	Created   string `json:"created"`
}

// Webhooks to model a list of webhooks
type Webhooks struct {
	Items []Webhook
}
