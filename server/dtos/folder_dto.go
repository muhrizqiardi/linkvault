package dtos

import (
	"github.com/google/uuid"
)

type CreateFolderDto struct {
	Name    string    `json:"name" db:"name" validation:"required,string"`
	OwnerId uuid.UUID `json:"owner_id" db:"owner_id" validation:"required,uuid" format:"uuid"`
}

type UpdateFolderDto struct {
	Name    string    `json:"name,omitempty" db:"name" validation:"required,string"`
	OwnerId uuid.UUID `json:"owner_id,omitempty" db:"owner_id" validation:"required,uuid" format:"uuid"`
}
