package model

import "database/sql"

type JadwalPsikolog struct {
	Hari         sql.NullString
	NamaPsikolog sql.NullString
	Spesialisasi sql.NullString
	JenisLayanan sql.NullString
	JamLayanan   sql.NullString
	Metode       sql.NullString
	Catatan      sql.NullString
}

// response for LLM
type JadwalPsikologResponse struct {
	Hari         string `json:"hari"`
	NamaPsikolog string `json:"nama_psikolog"`
	Spesialisasi string `json:"spesialisasi"`
	JenisLayanan string `json:"jenis_layanan"`
	JamLayanan   string `json:"jam_layanan"`
	Metode       string `json:"metode"`
	Catatan      string `json:"catatan"`
}
