package repository

import (
	"learn-golang/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(models.Reservation) (int, error)
	InsertRoomRestriction(models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID string) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomById(id int) (models.Room, error)
}
