package schemas

import "github.com/google/uuid"

type UserSchema struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	PassHash  []byte    `json:"pass_hash"`
	RoleID    uuid.UUID `json:"role_id"`
	SchoolID  uuid.UUID `json:"school_id"`
	ClassesID uuid.UUID `json:"classes_id"`
}
