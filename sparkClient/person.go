package sparkClient

import "time"

// SparkPerson represents a Cisco Spark Person
type SparkPerson struct {
	ID          string    `json:"id"`
	Emails      string    `json:"emails"`
	DisplayName string    `json:"displayName"`
	Avatar      string    `json:"avatar"`
	Created     time.Time `json:"created"`
}

// SparkPeople represents a slice of Cisco Spark Person(s)
type SparkPeople struct {
	Items []SparkPerson `json:"items"`
}

// List fetchs List people in your organization.
func (sp *SparkPeople) List() *SparkPeople {
	// Use sparkClient to fetch List people in your organization.

	return &SparkPeople{}
}

// GetByID fetchs a person by ID.
func (sp *SparkPeople) GetByID(ID string) *SparkPerson {

	return &SparkPerson{}
}

// GetMe fetchs the current user.
func (sp *SparkPeople) GetMe() *SparkPerson {

	return &SparkPerson{}
}
