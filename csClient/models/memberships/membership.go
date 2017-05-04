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

/*

   "id" : "Y2lzY29zcGFyazovL3VzL01FTUJFUlNISVAvMGQwYzkxYjYtY2U2MC00NzI1LWI2ZDAtMzQ1NWQ1ZDExZWYzOmNkZTFkZDQwLTJmMGQtMTFlNS1iYTljLTdiNjU1NmQyMjA3Yg",
   "roomId" : "Y2lzY29zcGFyazovL3VzL1JPT00vYmJjZWIxYWQtNDNmMS0zYjU4LTkxNDctZjE0YmIwYzRkMTU0",
   "personId" : "Y2lzY29zcGFyazovL3VzL1BFT1BMRS9mNWIzNjE4Ny1jOGRkLTQ3MjctOGIyZi1mOWM0NDdmMjkwNDY",
   "personEmail" : "r2d2@example.com",
   "isModerator" : true,
   "isMonitor" : true,
   "created" : "2015-10-18T14:26:16+00:00"

*/
