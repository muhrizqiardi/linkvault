package dtos

type CreateFolderDto struct {
	Name string `json:"name" db:"name" validation:"required,string"`
}

type UpdateFolderDto struct {
	Name string `json:"name,omitempty" db:"name" validation:"required,string"`
}
