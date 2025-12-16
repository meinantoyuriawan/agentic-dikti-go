package service

import (
	"AgenticDikti/internal/constants"
	"AgenticDikti/internal/model"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// load faq
type GeneralFAQ struct{}

func (w GeneralFAQ) Name() string {
	return "returning FAQ"
}

func (w GeneralFAQ) Description() string {
	return "returning FAQ answer."
}

func (w GeneralFAQ) Call(ctx context.Context, input string) (string, error) {
	content, err := os.ReadFile("faq.txt")

	return fmt.Sprintf("%d", content), err
}

// get jadwal psikolog
type PsychSchedule struct {
	service *Service
}

func (w PsychSchedule) Name() string {
	return "returning Psycologists schedule"
}

func (w PsychSchedule) Description() string {
	return `Creates a psychology booking.
	Input must be JSON:
	{
	  "Nama": "string",
	  "Nim": "string",
	  "Schedule": "string",
	  "UniversityID": "string"
	}`
}

func (w PsychSchedule) Call(ctx context.Context, input string) (string, error) {
	// function to call repository

	var booking model.BookingData

	if err := json.Unmarshal([]byte(input), &booking); err != nil {
		return "", fmt.Errorf("invalid input format: %w", err)
	}

	err := w.service.repository.InsertBooking(ctx, booking)
	if err != nil {
		return "", constants.ErrBooking
	}

	return `{"status":"success","message":"Booking created"}`, nil
}
