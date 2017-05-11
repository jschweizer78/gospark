package rooms

import "github.com/manyminds/api2go/jsonapi"

// Room is a model of a Spark Room
type Room struct {
	ID           string                                    `json:"id"`
	Title        string                                    `json:"title"`
	Type         string                                    `json:"type"`
	IsLocked     string                                    `json:"isLocked"`
	TeamID       string                                    `json:"teamId"`
	LastActivity string                                    `json:"lastActivity"`
	Created      string                                    `json:"created"`
	RefObjs      map[string][]jsonapi.Reference            `json:"-"`
	RefObjIDs    map[string]map[string]jsonapi.ReferenceID `json:"-"`
}

// Rooms to model a list of Rooms
type Rooms struct {
	Items []Room
}

// NewRoom is used to create an instance of a room
func NewRoom() *Room {
	return &Room{
		RefObjIDs: make(map[string]map[string]jsonapi.ReferenceID),
		RefObjs:   make(map[string][]jsonapi.Reference),
	}
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (rm *Room) GetID() string {
	return rm.ID
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (rm *Room) SetID(id string) error {
	rm.ID = id
	return nil
}

// SetToManyReferenceIDs sets the sweets reference IDs and satisfies the jsonapi.UnmarshalToManyRelations interface
func (rm *Room) SetToManyReferenceIDs(name string, IDs []string) error {
	switch name {
	default:
		for _, id := range IDs {
			rm.RefObjIDs[name][id] = jsonapi.ReferenceID{
				ID:   id,
				Type: "spark",
				Name: name,
			}
		}
		return nil
	}
}

// AddToManyIDs adds some new sweets that a users loves so much
func (rm *Room) AddToManyIDs(name string, IDs []string) error {
	// TODO add a check before setting the value
	switch name {
	default:
		for _, id := range IDs {
			rm.RefObjIDs[name][id] = jsonapi.ReferenceID{
				ID:   id,
				Type: "spark",
				Name: name,
			}
		}
		return nil
	}
}

// DeleteToManyIDs removes some sweets from a users because they made him very sick
func (rm *Room) DeleteToManyIDs(name string, IDs []string) error {
	switch name {
	default:
		for _, id := range IDs {
			delete(rm.RefObjIDs[name], id)
		}
		return nil
	}
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (rm *Room) GetReferences() []jsonapi.Reference {
	var lens []int64
	for _, val := range rm.RefObjs {
		lens = append(lens, int64(len(val)))
	}
	result := make([]jsonapi.Reference, sum(lens...))
	for _, cVal := range rm.RefObjs {
		for _, rVal := range cVal {
			result = append(result, rVal)
		}
	}
	return result
}

/*
// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (rm *Room) GetReferencedIDs() []jsonapi.ReferenceID {
	return rm.RefObjIDs
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (u User) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range u.Chocolates {
		result = append(result, u.Chocolates[key])
	}

	return result
}

*/

func sum(nums ...int64) int64 {
	var total int64
	for _, num := range nums {
		total += num
	}
	return total
}
