package people

// Person is a model of a "People"
type Person struct {
	ID            string   `json:"id"`
	Emails        string   `json:"emails"`
	DisplayName   string   `json:"displayName"`
	FirstName     string   `json:"firstName"`
	LastName      string   `json:"lastName"`
	Avatar        string   `json:"avatar"`
	OrgID         string   `json:"orgId"`
	Roles         []string `json:"roles"`
	Licenses      []string `json:"licenses"`
	Created       string   `json:"created"`
	Timezone      string   `json:"timezone"`
	LastActivity  string   `json:"lastActivity"`
	Status        string   `json:"status"`
	InvitePending bool     `json:"invitePending"`
	LoginEnabled  bool     `json:"loginEnabled"`
}

// People of a list of Persons
type People struct {
	Items Person `json:"items"`
}

/*
{
  "items" : [ {
    "id" : "Y2lzY29zcGFyazovL3VzL1BFT1BMRS9mNWIzNjE4Ny1jOGRkLTQ3MjctOGIyZi1mOWM0NDdmMjkwNDY",
    "emails" : [ "johnny.chang@foomail.com", "jchang@barmail.com" ],
    "displayName" : "John Andersen",
    "firstName" : "John",
    "lastName" : "Andersen",
    "avatar" : "https://1efa7a94ed21783e352-c62266528714497a17239ececf39e9e2.ssl.cf1.rackcdn.com/V1~54c844c89e678e5a7b16a306bc2897b9~wx29yGtlTpilEFlYzqPKag==~1600",
    "orgId" : "OTZhYmMyYWEtM2RjYy0xMWU1LWExNTItZmUzNDgxOWNkYzlh",
    "roles" : [ "Y2lzY29zcGFyazovL3VzL1JPT00vOGNkYzQwYzQtZjA5ZS0zY2JhLThjMjYtZGQwZTcwYWRlY2Iy", "Y2lzY29zcGFyazovL3VzL1BFT1BMRS9mMDZkNzFhNS0wODMzLTRmYTUtYTcyYS1jYzg5YjI1ZWVlMmX" ],
    "licenses" : [ "Y2lzY29zcGFyazovL3VzL1JPT00vOGNkYzQwYzQtZjA5ZS0zY2JhLThjMjYtZGQwZTcwYWRlY2Iy", "Y2lzY29zcGFyazovL3VzL1BFT1BMRS9mMDZkNzFhNS0wODMzLTRmYTUtYTcyYS1jYzg5YjI1ZWVlMmX" ],
    "created" : "2015-10-18T14:26:16+00:00",
    "timezone" : "America/Denver",
    "lastActivity" : "2015-10-18T14:26:16+00:00",
    "status" : "active",
    "invitePending" : false,
    "loginEnabled" : true
  } ]
}

*/
