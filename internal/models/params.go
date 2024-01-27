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
	Usename string `json:"username"`
	Password string `json:"password"`

}