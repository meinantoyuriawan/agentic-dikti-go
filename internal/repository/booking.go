package repository

import (
	"AgenticDikti/internal/model"
	"context"
)

const (
	insertBookingData = `INSERT INTO counseling_bookings (nama, nim, schedule, universityid) VALUES ($1, $2, $3, $4)`
)

func (q *Queries) InsertBooking(ctx context.Context, userBookData model.BookingData) (err error) {
	err = q.db.QueryRowContext(ctx, insertBookingData,
		userBookData.Nama,
		userBookData.Nim,
		userBookData.Schedule,
		userBookData.UniversityID,
	).Scan()

	if err != nil {
		return err
	}
	return nil
}
