package rooms

// Message is a model of a Spark Room
type Room struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Type         string `json:"type"`
	IsLocked     string `json:"isLocked"`
	TeamID       string `json:"teamId"`
	LastActivity string `json:"lastActivity"`
	Created      string `json:"created"`
}
