package dtos

type CreateLinkDto struct {
	Url      string `json:"url" db:"url" validate:"required,url"`
	Excerpt  string `json:"excerpt" db:"excerpt" validate:"required,string"`
	CoverUrl string `json:"cover_url" db:"cover_url" validate:"required,url"`
}

type UpdateLinkDto struct {
	Excerpt  string `json:"excerpt,omitempty" db:"excerpt" validate:"required,string"`
	CoverUrl string `json:"cover_url,omitempty" db:"cover_url" validate:"required,url"`
}
