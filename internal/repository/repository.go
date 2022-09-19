package repository

import (
	"learn-golang/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(models.Reservation) (int, error)
	InsertRoomRestriction(models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomById(int) (models.Room, error)
	GetUserById(int) (models.User, error)
	UpdateUser(models.User) error
	Authenticate(string, string) (int, string, error)
}
