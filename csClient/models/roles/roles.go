package roles

// Role is a model of a Spark role
type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Roles is a model of a list of roles
type Roles struct {
	Items []Role `json:"items"`
}
