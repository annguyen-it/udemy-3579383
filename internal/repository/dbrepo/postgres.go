package dbrepo

import (
	"context"
	"learn-golang/internal/models"
	"time"
)

func (*postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (rp *postgresDBRepo) InsertReservation(res models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO reservations 
                (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
             VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	_, err := rp.DB.ExecContext(
		ctx, stmt,
		res.FirstName, res.LastName, res.Email, res.Phone, res.StartDate, res.EndDate, res.RoomID,
		time.Now(), time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
