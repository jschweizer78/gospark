package rmSrv

import (
	"fmt"

	"github.com/jschweizer78/gospark/csClient/models/rooms"
)

type RmService struct {
	rmMap map[string]rooms.Room
}

// LoadRooms will load the rooms from a JSON return into the Room Service map
func (rs *RmService) LoadRooms(rms rooms.Rooms) {
	for _, rm := range rms.Items {
		rs.rmMap[rm.ID] = rm
	}

}

// GetRoomByID will return a room and a error indicating if it was founf
func (rs *RmService) GetRoomByID(id string) (rooms.Room, error) {
	item, found := rs.rmMap[id]
	if !found {
		return item, fmt.Errorf("the ID: %s was not found in map", id)
	}
	return item, nil
}
