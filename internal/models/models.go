package models

import(
	"github.com/google/uuid"
)

type DBUser struct{
	ID uuid.UUID
	Name string `json:"name"`
	Usename string `json:"username"`
}