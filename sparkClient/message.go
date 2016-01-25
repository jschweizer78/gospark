package sparkClient

import "time"

// SparkMessage represents a Cisco Spark Message
type SparkMessage struct {
	ID            string    `json:"id"`
	Emails        string    `json:"personId"`
	DisplayName   string    `json:"personEmail"`
	Avatar        string    `json:"roomId"`
	Text          string    `json:"text"`
	Files         string    `json:"files"`
	ToPersonID    string    `json:"toPersonId"`    //for private messages
	ToPersonEmail string    `json:"toPersonEmail"` //for private messages
	Created       time.Time `json:"created"`
}

// SparkMessages represents a slice of Cisco Spark Messag(s)
type SparkMessages struct {
	Items []SparkMessage `json:"items"`
}
