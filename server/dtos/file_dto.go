package dtos

import "github.com/google/uuid"

type CreateFileDto struct {
	LinkId  uuid.UUID `json:"link_id" db:"link_id" validation:"required,uuid" format:"uuid"`
	FileUrl string    `json:"file_url" db:"file_url" validation:"required,url"`
	OwnerId uuid.UUID `json:"owner_id" db:"owner_id" validation:"required,uuid" format:"uuid"`
}

type UpdateFileDto struct {
	LinkId  uuid.UUID `json:"link_id,omitempty" db:"link_id" validation:"required,uuid" format:"uuid"`
	FileUrl string    `json:"file_url,omitempty" db:"file_url" validation:"required,url"`
	OwnerId uuid.UUID `json:"owner_id,omitempty" db:"owner_id" validation:"required,uuid" format:"uuid"`
}
