package models

import(
	"github.com/google/uuid"
)

type CreateUserParams struct{
	ID uuid.UUID
	Name string `json:"name"`
	Usename string `json:"username"`
	Password string `json:"password"`
}

type LoginUserParams struct{
	ID uuid.UUID	`json:"id,omitempty"`
	Usename string `json:"username"`
	Password string `json:"password"`

}

type CreatePostParams struct {
	ImageURL string `json:"image_url"`
	Description string `json:"description"`
}