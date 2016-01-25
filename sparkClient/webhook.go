package sparkClient

import "time"

// SparkWebhook represents a Cisco Spark Webhook
type SparkWebhook struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	TargetURL string    `json:"targetUrl"`
	Resource  string    `json:"resource"`
	Event     string    `json:"event"`
	Filter    string    `json:"filter"`
	Create    time.Time `json:"created"`
}

// SparkWebhooks represents a slice of Cisco Spark Webhook(s)
type SparkWebhooks struct {
	Items []SparkMessage `json:"items"`
}
