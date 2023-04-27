package dtos

import "github.com/google/uuid"

type CreateTagDto struct {
	Name    string    `json:"name" db:"name" validation:"required,name"`
	LinkId  uuid.UUID `json:"link_id" db:"link_id" validation:"required"`
	OwnerId uuid.UUID `json:"owner_id" db:"owner_id" validation:"required,uuid" format:"uuid"`
}

type UpdateTagDto struct {
	Name    string    `json:"name,omitempty" db:"name" validation:"required,name"`
	LinkId  uuid.UUID `json:"link_id,omitempty" db:"link_id" validation:"required"`
	OwnerId uuid.UUID `json:"owner_id,omitempty" db:"owner_id" validation:"required,uuid" format:"uuid"`
}
