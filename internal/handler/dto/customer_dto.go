package dto

type CreateConsumerRequest struct {
	NIK            string  `json:"nik" validate:"required"`
	FullName       string  `json:"full_name" validate:"required"`
	LegalName      string  `json:"legal_name" validate:"required"`
	BirthPlace     string  `json:"birth_place"`
	BirthDate      string  `json:"birth_date"` // "YYYY-MM-DD"
	Salary         float64 `json:"salary"`
	KTPPhotoURL    string  `json:"ktp_photo_url"`
	SelfiePhotoURL string  `json:"selfie_photo_url"`
}

type ConsumerResponse struct {
	ID             string  `json:"id"`
	NIK            string  `json:"nik"`
	FullName       string  `json:"full_name"`
	LegalName      string  `json:"legal_name"`
	BirthPlace     string  `json:"birth_place"`
	BirthDate      string  `json:"birth_date"`
	Salary         float64 `json:"salary"`
	KTPPhotoURL    string  `json:"ktp_photo_url"`
	SelfiePhotoURL string  `json:"selfie_photo_url"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
}
