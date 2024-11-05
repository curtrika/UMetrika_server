package schemas

import "github.com/google/uuid"

type UserSchema struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	PassHash []byte    `json:"pass_hash"`
}
