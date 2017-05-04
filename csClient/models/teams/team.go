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

// Teams to model a list of Team
type Teams struct {
	Items []Teams
}

// TeamMemberships to model a list of TeamMembership
type TeamMemberships struct {
	Items []TeamMembership
}
