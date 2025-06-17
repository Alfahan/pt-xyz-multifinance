package domain

import "time"

type Consumer struct {
	ID             string    `json:"id"`
	NIK            string    `json:"nik"`
	FullName       string    `json:"full_name"`
	LegalName      string    `json:"legal_name"`
	BirthPlace     string    `json:"birth_place"`
	BirthDate      time.Time `json:"birth_date"`
	Salary         float64   `json:"salary"`
	KTPPhotoURL    string    `json:"ktp_photo_url"`
	SelfiePhotoURL string    `json:"selfie_photo_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
