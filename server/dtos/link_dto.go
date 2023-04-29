package dtos

type CreateLinkDto struct {
	Url      string `json:"url" db:"url" validate:"required,url" format:"url"`
	Title    string `json:"title" db:"title" validation:"required"`
	Excerpt  string `json:"excerpt" db:"excerpt" validate:"required"`
	CoverUrl string `json:"cover_url" db:"cover_url" validate:"required,url" format:"url"`
}

type UpdateLinkDto struct {
	Title    string `json:"title,omitempty" db:"title" validation:"required"`
	Excerpt  string `json:"excerpt,omitempty" db:"excerpt" validate:"required" format:"url"`
	CoverUrl string `json:"cover_url,omitempty" db:"cover_url" validate:"required,url" format:"url"`
}
