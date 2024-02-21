package model

import "mime/multipart"

type RequestDonation struct {
	Email string                `json:"email"`
	Photo *multipart.FileHeader `form:"photo"`
}
