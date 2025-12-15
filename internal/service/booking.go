package service

import (
	"AgenticDikti/internal/constants"
	"AgenticDikti/internal/model"
	"context"
)

func (s *Service) InputBooking(ctx context.Context, userBookData model.BookingData) (err error) {
	err = s.repository.InsertBooking(ctx, userBookData)
	if err != nil {
		return constants.ErrBooking
	}

	return nil
}
