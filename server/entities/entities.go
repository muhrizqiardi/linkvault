package entities

import (
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	Id        uuid.UUID `json:"id" db:"id" validation:"required,uuid" format:"uuid"`
	Email     string    `json:"email" db:"email" validation:"required,email" format:"email" `
	FullName  string    `json:"full_name" db:"full_name" validation:"required,string"`
	Password  string    `json:"-" db:"password" validation:"required,string,min=8" minLength:"8"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validation:"date" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validation:"date" format:"date-time"`
}

type FolderEntity struct {
	Id        uuid.UUID `json:"id" db:"id" validation:"required,uuid" format:"uuid"`
	Name      string    `json:"name" db:"name" validation:"required,string"`
	OwnerId   uuid.UUID `json:"owner_id" db:"owner_id" validation:"required,uuid" format:"uuid"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validation:"date" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validation:"date" format:"date-time"`
}

type TagEntity struct {
	Id        uuid.UUID `json:"id" db:"id" validation:"required,uuid" format:"uuid"`
	Name      string    `json:"name" db:"name" validation:"required,name"`
	OwnerId   uuid.UUID `json:"owner_id" db:"owner_id" validation:"required,uuid" format:"uuid"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validation:"date" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validation:"date" format:"date-time"`
}

type LinkEntity struct {
	Id        uuid.UUID `json:"id" db:"id" validation:"required,uuid" format:"uuid"`
	Url       string    `json:"url" db:"url" validation:"required,url"`
	Excerpt   string    `json:"excerpt" db:"excerpt" validation:"required,string"`
	CoverUrl  string    `json:"cover_url" db:"cover_url" validation:"required,url"`
	OwnerId   uuid.UUID `json:"owner_id" db:"owner_id" validation:"required,uuid" format:"uuid"`
	FolderId  uuid.UUID `json:"folder_id" db:"folder_id" validation:"required,uuid" format:"uuid"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validation:"date" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validation:"date" format:"date-time"`
}

type LinkMediaEntity struct {
	Id        uuid.UUID `json:"id" db:"id" validation:"required,uuid" format:"uuid"`
	LinkId    uuid.UUID `json:"link_id" db:"link_id" validation:"required,uuid" format:"uuid"`
	MediaUrl  string    `json:"media_url" db:"media_url" validation:"required,url"`
	OwnerId   uuid.UUID `json:"owner_id" db:"owner_id" validation:"required,uuid" format:"uuid"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validation:"date" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validation:"date" format:"date-time"`
}

type FileEntity struct {
	Id        uuid.UUID `json:"id" db:"id" validation:"required,uuid" format:"uuid"`
	LinkId    uuid.UUID `json:"link_id" db:"link_id" validation:"required,uuid" format:"uuid"`
	FileUrl   string    `json:"file_url" db:"file_url" validation:"required,url"`
	OwnerId   uuid.UUID `json:"owner_id" db:"owner_id" validation:"required,uuid" format:"uuid"`
	CreatedAt time.Time `json:"created_at" db:"created_at" validation:"date" format:"date-time"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" validation:"date" format:"date-time"`
}
