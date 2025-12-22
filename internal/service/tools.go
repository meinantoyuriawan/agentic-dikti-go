package service

import (
	"AgenticDikti/internal/constants"
	"AgenticDikti/internal/model"
	"AgenticDikti/internal/utils"
	"context"
	"encoding/json"
	"fmt"
	"os"
)

// load faq
type GeneralFAQ struct{}

func (w GeneralFAQ) Name() string {
	return "generalFAQ"
}

func (w GeneralFAQ) Description() string {
	return "Use this tool ONLY when the user asks for tips, methods, or guidance on handling bullying situations. Returns psychological first aid advice."
}

func (w GeneralFAQ) Call(ctx context.Context, input string) (string, error) {
	content, err := os.ReadFile("internal/service/faq.txt")
	if err != nil {
		return "", err
	}

	// Return clean text, NOT bytes or numbers
	return string(content), nil
}

// get jadwal psikolog
type PsychSchedule struct {
	service *Service
}

func (w PsychSchedule) Name() string {
	return "jadwalPsikolog"
}

func (w PsychSchedule) Description() string {
	return "Use this tool ONLY when the user asks for campus psychologist schedules or agrees to see a psychologist. Returns available schedules."
}

func (w PsychSchedule) Call(ctx context.Context, input string) (string, error) {
	// function to call repository

	fmt.Println("schedule psy called")
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

func ToJadwalPsikologResponse(data []model.JadwalPsikolog) []model.JadwalPsikologResponse {

	result := make([]model.JadwalPsikologResponse, 0, len(data))

	for _, d := range data {
		result = append(result, model.JadwalPsikologResponse{
			Hari:         utils.NullStringToString(d.Hari),
			NamaPsikolog: utils.NullStringToString(d.NamaPsikolog),
			Spesialisasi: utils.NullStringToString(d.Spesialisasi),
			JenisLayanan: utils.NullStringToString(d.JenisLayanan),
			JamLayanan:   utils.NullStringToString(d.JamLayanan),
			Metode:       utils.NullStringToString(d.Metode),
			Catatan:      utils.NullStringToString(d.Catatan),
		})
	}

	return result
}

// book psikolog
type BookPsychologist struct {
	service *Service
}

func (w BookPsychologist) Name() string {
	return "bookingPsikolog"
}

func (w BookPsychologist) Description() string {
	return "Create a psychologist booking AFTER the user confirms name, NIM, and schedule. Input must be valid JSON."
}

func (w BookPsychologist) Call(ctx context.Context, input string) (string, error) {
	// function to call repository

	var booking model.BookingData
	fmt.Println("book psy called")

	if err := json.Unmarshal([]byte(input), &booking); err != nil {
		return "", fmt.Errorf("invalid input format: %w", err)
	}

	err := w.service.repository.InsertBooking(ctx, booking)
	if err != nil {
		return "", constants.ErrBooking
	}

	result := map[string]string{
		"status":  "success",
		"message": "Booking psikolog berhasil dibuat",
	}

	bytes, _ := json.Marshal(result)
	return string(bytes), nil
}
