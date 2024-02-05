package models

import(
	"github.com/google/uuid"
)

type DBUser struct{
	ID uuid.UUID
	Name string `json:"name"`
	Usename string `json:"username"`
}

type DBPost struct{
	ID uuid.UUID `json:"id"`
	UserID uuid.UUID `json:"user_id"`
	ImageURL string `json:"image_url"`
	Description string `json:"description"`
} 

type DBFollow struct{
	FollowerID uuid.UUID
	FollowingID uuid.UUID
}

type DBLike struct{
	PostID uuid.UUID
	UserID uuid.UUID
}