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
func (rp *postgresDBRepo) InsertReservation(m models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newId int
	stmt := `
        INSERT INTO reservations 
            (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id
    `

	err := rp.DB.QueryRowContext(
		ctx, stmt,
		m.FirstName, m.LastName, m.Email, m.Phone, m.StartDate, m.EndDate, m.RoomID,
		time.Now(), time.Now(),
	).Scan(&newId)

	if err != nil {
		return 0, err
	}

	return newId, nil
}

// InsertRoomRestriction inserts a new room restriction into the database
func (rp *postgresDBRepo) InsertRoomRestriction(m models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
        INSERT INTO room_restrictions 
            (start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

	_, err := rp.DB.ExecContext(
		ctx, stmt,
		m.StartDate, m.EndDate, m.RoomID, m.ReservationID, time.Now(), time.Now(), m.RestrictionID,
	)

	if err != nil {
		return err
	}

	return nil
}

// SearchAvailabilityByDates returns true if availability exists for roomID, and false otherwise
func (rp *postgresDBRepo) SearchAvailabilityByDates(start, end time.Time, roomID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int

	query := `
        SELECT COUNT(id) 
        FROM room_restrictions 
        WHERE 
            roomID = $1
            $2 < end_date AND $3 > start_date
    `

	row := rp.DB.QueryRowContext(
		ctx, query,
		roomID, start, end,
	)
	err := row.Scan(&numRows)
	if err != nil {
		return false, err
	}

	return numRows == 0, nil
}
