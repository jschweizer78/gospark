package teams

// Team is a model of a Spark Team
type Team struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
}

// TeamMembership is a model of a Spark Team Membership
type TeamMembership struct {
	ID          string `json:"id"`
	TeamID      string `json:"teamId"`
	PersonID    string `json:"personId"`
	PersonEmail string `json:"personEmailid"`
	IsModerator bool   `json:"isModerator"`
	Created     string `json:"created"`
}
