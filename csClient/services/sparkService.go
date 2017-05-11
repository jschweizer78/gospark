package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jschweizer78/gospark/csClient"
	"github.com/jschweizer78/gospark/csClient/models/people"
	"github.com/jschweizer78/gospark/csClient/models/rooms"
)

var sparkURL = "https://api.ciscospark.com/v1"

// SparkServ is used to CRUD the Spark Cloud.
type SparkServ struct {
	cs *csClient.SparkClient
}

// NewService is used to create new Spark service
func NewService(cs *csClient.SparkClient) (*SparkServ, error) {
	return &SparkServ{cs}, nil
}

func (ss *SparkServ) ProcessRequest(req *http.Request) (response []byte) {
	return ss.processRequest(req)
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

// NewRoom is used to create a new Spark room
// This function will create a new Spark Room
func (ss *SparkServ) NewRoom(roomName string) rooms.Room {
	jsonString := fmt.Sprintf("{ \"title\" : \"%s\" }", roomName)
	var jsonStr = []byte(jsonString)
	req, err := http.NewRequest("POST", sparkURL+"/rooms", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("creating request failed:", err)
	}
	jsonOut := ss.processRequest(req)
	var room rooms.Room
	err = json.Unmarshal(jsonOut, &room)
	if err != nil {
		log.Println("error: ", err)
	}
	return room
}

// ListRooms Gets a list of Room for a given Auth Token
// This function lists all the rooms in the Spark api.
func (ss *SparkServ) ListRooms() rooms.Rooms {
	req, err := http.NewRequest("GET", sparkURL+"/rooms", nil)
	if err != nil {
		log.Println("creating request failed:", err)
	}

	jsonOut := ss.processRequest(req)
	var rooms rooms.Rooms
	err = json.Unmarshal(jsonOut, &rooms)
	if err != nil {
		log.Println("error: ", err)
	}
	//fmt.Printf("%+v", rooms)
	return rooms
}

// GetRoomByID Gets a of Room for by ID
// This function lists all the rooms in the Spark api.
func (ss *SparkServ) GetRoomByID(id string) rooms.Room {
	reqParam := fmt.Sprintf("%s/rooms/%s", sparkURL, id)
	req, err := http.NewRequest("GET", reqParam, nil)
	if err != nil {
		log.Println("creating request failed:", err)
	}

	jsonOut := ss.processRequest(req)
	var room rooms.Room
	err = json.Unmarshal(jsonOut, &room)
	if err != nil {
		log.Println("error: ", err)
	}
	//fmt.Printf("%+v", rooms)
	return room
}
