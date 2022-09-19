package dbrepo

import (
	"errors"
	"learn-golang/internal/models"
	"time"
)

func (*testDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (rp *testDBRepo) InsertReservation(_ models.Reservation) (int, error) {
	return 1, nil
}

// InsertRoomRestriction inserts a new room restriction into the database
func (rp *testDBRepo) InsertRoomRestriction(_ models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByDatesByRoomID returns true if availability exists for roomID, and false otherwise
func (rp *testDBRepo) SearchAvailabilityByDatesByRoomID(_, _ time.Time, _ int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns a slice of available rooms, if any, for given date range
func (rp *testDBRepo) SearchAvailabilityForAllRooms(_, _ time.Time) (rooms []models.Room, err error) {
	return
}

// GetRoomById gets a room by id
func (rp *testDBRepo) GetRoomById(id int) (models.Room, error) {
	var room models.Room
	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil
}

func (rp *testDBRepo) GetUserById(_ int) (models.User, error) {
	var u models.User
	return u, nil
}

func (rp *testDBRepo) UpdateUser(_ models.User) error {
	return nil
}

func (rp *testDBRepo) Authenticate(_, _ string) (int, string, error) {
	return 0, "", nil
}
