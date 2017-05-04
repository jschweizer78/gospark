package messages

// Message is a model of a message recv'd
type Message struct {
	ID              string   `json:"id"`
	RoomID          string   `json:"roomId"`
	RoomType        string   `json:"roomType"`
	ToPersonID      string   `json:"toPersonId"`
	ToPersonEmail   string   `json:"toPersonEmail"`
	Text            string   `json:"text"`
	Markdown        string   `json:"markdown"`
	Files           []string `json:"files"`
	PersonID        string   `json:"personId"`
	PersonEmail     string   `json:"personEmail"`
	Created         string   `json:"created"`
	MentionedPeople []string `json:"mentionedPeople"`
}
