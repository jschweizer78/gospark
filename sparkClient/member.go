package sparkClient

import "time"

// SparkMember represents a Cisco Spark Person
type SparkMember struct {
	ID          string    `json:"id"`
	Emails      string    `json:"emails"`
	DisplayName string    `json:"displayName"`
	Avatar      string    `json:"avatar"`
	Created     time.Time `json:"created"`
}

// SparkMembers represents a slice of Cisco Spark Person(s)
type SparkMembers struct {
	Items []SparkMember `json:"items"`
}
