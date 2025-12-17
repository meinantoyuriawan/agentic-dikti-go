package service

import (
	"AgenticDikti/internal/constants"
	"AgenticDikti/internal/model"
	"context"
	"database/sql"
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
	return `Get a Psycologists schedule.`
}

func (w PsychSchedule) Call(ctx context.Context, input string) (string, error) {
	// function to call repository

	schedules, err := w.service.repository.SelectJadwalPsikolog(ctx)
	if err != nil {
		return "", constants.ErrBooking
	}

	schedulesString := ToJadwalPsikologResponse(schedules)

	bytes, err := json.Marshal(schedulesString)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func ToJadwalPsikologResponse(data []model.JadwalPsikolog) []model.JadwalPsikologResponse {

	result := make([]model.JadwalPsikologResponse, 0, len(data))

	for _, d := range data {
		result = append(result, model.JadwalPsikologResponse{
			Hari:         nullStringToString(d.Hari),
			NamaPsikolog: nullStringToString(d.NamaPsikolog),
			Spesialisasi: nullStringToString(d.Spesialisasi),
			JenisLayanan: nullStringToString(d.JenisLayanan),
			JamLayanan:   nullStringToString(d.JamLayanan),
			Metode:       nullStringToString(d.Metode),
			Catatan:      nullStringToString(d.Catatan),
		})
	}

	return result
}

// book psikolog
type BookPsychologist struct {
	service *Service
}

func (w BookPsychologist) Name() string {
	return "creates a Psycologists booking"
}

func (w BookPsychologist) Description() string {
	return `Creates a psychology booking.
	Input must be JSON:
	{
		"Nama": "string",
		"Nim": "string",
		"Schedule": "string",
		"UniversityID": "string"
	}`
}

func (w BookPsychologist) Call(ctx context.Context, input string) (string, error) {
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
