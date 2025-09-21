package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
)

type Attachment struct {
	orm.Model
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Deadline       string
	AttachmentName string
	ReferenceID    string
	ReferenceName  string
	FileName       string
	StoragePath    string
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
	DeletedAt      *time.Time
	CreatedBy      *string
	UpdatedBy      *string
	DeletedBy      *string
}
