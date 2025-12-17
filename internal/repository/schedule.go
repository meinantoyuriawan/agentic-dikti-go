package repository

import (
	"AgenticDikti/internal/model"
	"context"
)

const (
	selectJadwalPsikologQuery = `SELECT jp.hari, jp.nama_psikolog, jp.spesialisasi, jp.jenis_layanan, jp.jam_layanan, jp.metode, jp.catatan FROM jadwal_psikolog jp`
)

func (q *Queries) SelectJadwalPsikolog(ctx context.Context) (res []model.JadwalPsikolog, err error) {
	rows, err := q.db.QueryContext(ctx, selectJadwalPsikologQuery)
	// Scan(&res.Role, &res.ChatInput)
	if err != nil {
		return []model.JadwalPsikolog{}, err
	}
	defer rows.Close()

	var schedules []model.JadwalPsikolog

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var schedule model.JadwalPsikolog
		if err := rows.Scan(
			&schedule.Hari,
			&schedule.NamaPsikolog,
			&schedule.Spesialisasi,
			&schedule.JenisLayanan,
			&schedule.JamLayanan,
			&schedule.Metode,
			&schedule.Catatan); err != nil {
			return schedules, err
		}
		schedules = append(schedules, schedule)
	}
	if err = rows.Err(); err != nil {
		return schedules, err
	}
	return schedules, nil
}
