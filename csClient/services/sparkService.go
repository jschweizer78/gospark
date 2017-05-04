package services

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jschweizer78/gospark/csClient"
	"github.com/jschweizer78/gospark/csClient/models/people"
	"github.com/jschweizer78/gospark/csClient/models/rooms"
)

var sparkURL string = "https://api.ciscospark.com/v1"

type SparkServ struct {
	cs *csClient.SparkClient
}

// NewService is used to create new Spark service
func NewService(cs *csClient.SparkClient) (*SparkServ, error) {
	return &SparkServ{cs}, nil
}

func (ss *SparkServ) processRequest(req *http.Request) (response []byte) {
	req.Header.Set("Authorization", "Bearer "+ss.cs.Authtoken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := ss.cs.HttpClient.Do(req)
	if err != nil {
		log.Println("Error: ", err)
		//log.Println("Error: ", resp.StatusCode)
		return
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("StatusCode =", resp.StatusCode)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	// log.Println("StatusCode: ", resp.StatusCode)
	//fmt.Printf("%s", body)
	return body
}

// PostToRoom is used to Pose a message to a Spark Room
func (ss *SparkServ) PostToRoom(text string, room rooms.Room, fileURL string) {
	var jsonString string
	if fileURL != "" {
		jsonString = fmt.Sprintf("{ \"roomId\" : \"%s\" ,  \"file\" : \"%s\" , \"text\" : \"%s\" }", room.ID, fileURL, text)
	} else {
		jsonString = fmt.Sprintf("{ \"roomId\" : \"%s\" ,  \"text\" : \"%s\" }", room.ID, text)
	}
	//fmt.Println(jsonString)
	var jsonStr = []byte(jsonString)
	req, err := http.NewRequest("POST", sparkURL+"/messages", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("creating request failed:", err)
	}

	// ss.processRequest(req)
	resp := ss.processRequest(req)
	fmt.Printf("%s", resp)
}

// AddMemberToRoom is used to add a new person to a Spark Room
func (ss *SparkServ) AddMemberToRoom(mem people.Person, room rooms.Room, isModerator bool) {
	jsonString := fmt.Sprintf("{ \"roomId\" : \"%s\" , \"personEmail\" : \"%s\", \"isModerator\" : %t }", room.ID, mem.Emails[0], isModerator)
	//fmt.Println(jsonString)
	var jsonStr = []byte(jsonString)
	req, err := http.NewRequest("POST", sparkURL+"/memberships", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("creating request failed:", err)
	}
	ss.processRequest(req)
	//str := s.processRequest(req)
	//fmt.Printf("%s", str)
}
