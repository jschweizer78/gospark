// Cisco Spark Client to get information from the spark service.
//
package csClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type SparkRoom struct {
	Id           string
	Title        string
	LastActivity time.Time
	Create       time.Time
}

type SparkRooms struct {
	Items []SparkRoom
}

type SparkClient struct {
	authtoken  string
	conn       net.Conn
	httpClient *http.Client
	reader     io.ReadCloser
}

var sparkURL string = "https://api.ciscospark.com/v1"

// This function is used by the client requests to perform the transaction
func (s *SparkClient) processRequest(req *http.Request) (response []byte) {
	req.Header.Set("Authorization", "Bearer "+s.authtoken)
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.httpClient.Do(req)
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

func (s *SparkClient) UpdateAuth(auth string) {
	s.authtoken = auth
}

func (s *SparkClient) PostMessageToSparkRoom(text string, roomId string, fileURL string) {
	var jsonString string
	if fileURL != "" {
		jsonString = fmt.Sprintf("{ \"roomId\" : \"%s\" ,  \"file\" : \"%s\" , \"text\" : \"%s\" }", roomId, fileURL, text)
	} else {
		jsonString = fmt.Sprintf("{ \"roomId\" : \"%s\" ,  \"text\" : \"%s\" }", roomId, text)
	}
	//fmt.Println(jsonString)
	var jsonStr = []byte(jsonString)
	req, err := http.NewRequest("POST", sparkURL+"/messages", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("creating request failed:", err)
	}

	s.processRequest(req)
	//str := s.processRequest(req)
	//fmt.Printf("%s", str)
}

func (s *SparkClient) AddMemberToSparkRoom(email string, roomId string, isModerator bool) {
	jsonString := fmt.Sprintf("{ \"roomId\" : \"%s\" , \"personEmail\" : \"%s\", \"isModerator\" : %t }", roomId, email, isModerator)
	//fmt.Println(jsonString)
	var jsonStr = []byte(jsonString)
	req, err := http.NewRequest("POST", sparkURL+"/memberships", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("creating request failed:", err)
	}
	s.processRequest(req)
	//str := s.processRequest(req)
	//fmt.Printf("%s", str)
}

// This function will create a new Spark Room
func (s *SparkClient) NewRoom(roomName string) SparkRoom {
	jsonString := fmt.Sprintf("{ \"title\" : \"%s\" }", roomName)
	var jsonStr = []byte(jsonString)
	req, err := http.NewRequest("POST", sparkURL+"/rooms", bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println("creating request failed:", err)
	}
	jsonOut := s.processRequest(req)
	var room SparkRoom
	err = json.Unmarshal(jsonOut, &room)
	if err != nil {
		log.Println("error: ", err)
	}
	return room
}

// This function lists all the rooms in the Spark api.
func (s *SparkClient) Rooms() []SparkRoom {
	req, err := http.NewRequest("GET", sparkURL+"/rooms", nil)
	if err != nil {
		log.Println("creating request failed:", err)
	}

	jsonOut := s.processRequest(req)
	var rooms SparkRooms
	err = json.Unmarshal(jsonOut, &rooms)
	if err != nil {
		log.Println("error: ", err)
	}
	//fmt.Printf("%+v", rooms)
	return rooms.Items
}

// New Creates a new SparkClient to be used.
// s := SparkClient.New()
// s.Rooms()
// requires that the environment variable: "SPARK_AUTH_TOKEN" be defined
func New() *SparkClient {
	var conn net.Conn
	var r io.ReadCloser

	var conf struct {
		AuthToken string
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				if conn != nil {
					conn.Close()
					conn = nil
				}
				netc, err := net.DialTimeout(netw, addr, 5*time.Second)
				if err != nil {
					return nil, err
				}
				conn = netc
				return netc, nil
			},
		},
	}

	return &SparkClient{
		authtoken:  conf.AuthToken,
		conn:       conn,
		httpClient: client,
		reader:     r,
	}
}
