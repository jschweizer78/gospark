package sparkClient

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// SparkPerson represents a Cisco Spark Person
type SparkPerson struct {
	ID          string    `json:"id"`
	Emails      string    `json:"emails"`
	DisplayName string    `json:"displayName"`
	Avatar      string    `json:"avatar"`
	Created     time.Time `json:"created"`
	sc          *SparkClient
}

// SparkPeople represents a slice of Cisco Spark Person(s)
type SparkPeople struct {
	Items []SparkPerson `json:"items"`
	sc    *SparkClient
}

// PeopleQueryParam represents a query to the people list service
type PeopleQueryParam struct {
	Email       string
	DisplayName string
	max         int
}

// List fetchs List people in your organization.
func (sp *SparkPeople) List() *SparkPeople {
	// Use sparkClient to fetch List people in your organization.
	req, err := http.NewRequest("GET", sparkURL+"/people", nil)
	if err != nil {
		log.Println("creating request failed:", err)
	}

	jsonOut := sp.sc.processRequest(req)
	var people SparkPeople
	err = json.Unmarshal(jsonOut, &people)
	if err != nil {
		log.Println("error: ", err)
	}
	//fmt.Printf("%+v", rooms)
	return &people

}

// GetByID fetchs a person by ID.
func (sp *SparkPeople) GetByID(ID string) *SparkPerson {
	// Use sparkClient to fetch List people in your organization.
	req, err := http.NewRequest("GET", sparkURL+"/people/personID="+ID, nil)
	if err != nil {
		log.Println("creating request failed:", err)
	}

	jsonOut := sp.sc.processRequest(req)
	var person SparkPerson
	err = json.Unmarshal(jsonOut, &person)
	if err != nil {
		log.Println("error: ", err)
	}
	//fmt.Printf("%+v", rooms)
	return &person
}

// GetMe fetchs the current user.
func (sp *SparkPerson) GetMe() *SparkPerson {
	// Use sparkClient to fetch List people in your organization.
	req, err := http.NewRequest("GET", sparkURL+"/people/me", nil)
	if err != nil {
		log.Println("creating request failed:", err)
	}

	jsonOut := sp.sc.processRequest(req)
	var person SparkPerson
	err = json.Unmarshal(jsonOut, &person)
	if err != nil {
		log.Println("error: ", err)
	}
	return &person
}
