package dtos

import "github.com/google/uuid"

type CreateLinkDto struct {
	Url      string    `json:"url" db:"url" validate:"required,url"`
	Excerpt  string    `json:"excerpt" db:"excerpt" validate:"required,string"`
	CoverUrl string    `json:"cover_url" db:"cover_url" validate:"required,url"`
	OwnerId  uuid.UUID `json:"owner_id" db:"owner_id" validate:"required,uuid"`
	FolderId uuid.UUID `json:"folder_id" db:"folder_id" validate:"required,uuid"`
}

type UpdateLinkDto struct {
	Excerpt  string    `json:"excerpt,omitempty" db:"excerpt" validate:"required,string"`
	CoverUrl string    `json:"cover_url,omitempty" db:"cover_url" validate:"required,url"`
	OwnerId  uuid.UUID `json:"owner_id,omitempty" db:"owner_id" validate:"required,uuid"`
	FolderId uuid.UUID `json:"folder_id,omitempty" db:"folder_id" validate:"required,uuid"`
}
