package dtos

import "github.com/google/uuid"

type CreateLinkMediaDto struct {
	LinkId   uuid.UUID `json:"link_id" db:"link_id" validation:"required,uuid" format:"uuid"`
	MediaUrl string    `json:"media_url" db:"media_url" validation:"required,url"`
	OwnerId  uuid.UUID `json:"owner_id" db:"owner_id" validation:"required,uuid" format:"uuid"`
}

type UpdateLinkMediaDto struct {
	LinkId   uuid.UUID `json:"link_id,omitempty" db:"link_id" validation:"required,uuid" format:"uuid"`
	MediaUrl string    `json:"media_url,omitempty" db:"media_url" validation:"required,url"`
	OwnerId  uuid.UUID `json:"owner_id,omitempty" db:"owner_id" validation:"required,uuid" format:"uuid"`
}
