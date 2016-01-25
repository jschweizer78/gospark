package sparkClient

import "time"

// SparkRoom represents a Cisco Spark Room (probably red as the dir is called that.. should just be called rooms)
type SparkRoom struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	SipAddress   string `json:"sipAddress"`
	LastActivity time.Time
	Created      time.Time `json:"created"`
}

// SparkRooms represents a slice of Cisco Spark Rooms @ "Items"
type SparkRooms struct {
	Items []SparkRoom `json:"items"`
}
