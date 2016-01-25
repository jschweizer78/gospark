package sparkClient

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"
)

// SparkRoom represents a Cisco Spark Room (probably red as the dir is called that.. should just be called rooms)
type SparkRoom struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	SipAddress   string `json:"sipAddress"`
	LastActivity time.Time
	Created      time.Time `json:"created"`
	sc           *SparkClient
}

// SparkRooms represents a slice of Cisco Spark Rooms @ "Items"
type SparkRooms struct {
	Items []SparkRoom `json:"items"`
}

// PostMessage used to post to this Cisco Spark Room.
func (sr *SparkRoom) PostMessage(text string, fileURL string) {
	var jsonString string
	if fileURL != "" {
		jsonString = fmt.Sprintf("{ \"roomId\" : \"%s\" ,  \"file\" : \"%s\" , \"text\" : \"%s\" }", sr.ID, fileURL, text)
	} else {
		jsonString = fmt.Sprintf("{ \"roomId\" : \"%s\" ,  \"text\" : \"%s\" }", sr.ID, text)
	}
	//fmt.Println(jsonString)
	var jsonStr = []byte(jsonString)
	req, err := http.NewRequest("POST", sparkURL+"/messages", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("creating request failed:", err)
	}

	sr.sc.processRequest(req)
	// need to add response processing as depicted below
	//str := s.processRequest(req)
	//fmt.Printf("%s", str)
}

// AddMember used to add people to a Cisco Spark Room(this will be added as a method of the Room struct)
func (sr *SparkRoom) AddMember(email string, isModerator bool) {
	jsonString := fmt.Sprintf("{ \"roomId\" : \"%s\" , \"personEmail\" : \"%s\", \"isModerator\" : %t }", sr.ID, email, isModerator)
	//fmt.Println(jsonString)
	var jsonStr = []byte(jsonString)
	req, err := http.NewRequest("POST", sparkURL+"/memberships", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("creating request failed:", err)
	}
	sr.sc.processRequest(req)
	//str := s.processRequest(req)
	//fmt.Printf("%s", str)
}
