package memberships

// Membership models a Membership item
type Membership struct {
	ID          string `json:"id"`
	RoomID      string `json:"roomId"`
	PersonID    string `json:"personId"`
	PersonEmail string `json:"personEmail"`
	IsModerator bool   `json:"isModerator"`
	IsMonitor   bool   `json:"isMonitor"`
	Created     string `json:"created"`
}

// Memberships to model a list of Memberships
type Memberships struct {
	Items []Membership
}
