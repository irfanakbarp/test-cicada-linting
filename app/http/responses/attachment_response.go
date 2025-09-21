package responses

import (
	"goravel/app/models"
	"time"

	"github.com/google/uuid"
)

// This file used to map the Attachment model into a structured JSON response
// and ensuring field order and format consistency.

type AttachmentResponse struct {
	ID             uuid.UUID  `json:"id"`
	AttachmentName string     `json:"attachment_name"`
	FileName       string     `json:"file_name"`
	StoragePath    string     `json:"storage_path"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

func NewAttachmentResponse(resp models.Attachment) AttachmentResponse {
	return AttachmentResponse{
		ID:             resp.ID,
		AttachmentName: resp.AttachmentName,
		FileName:       resp.FileName,
		StoragePath:    resp.StoragePath,
		CreatedAt:      resp.CreatedAt,
		UpdatedAt:      resp.UpdatedAt,
	}
}
