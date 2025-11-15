package dtos

type UserDto struct {
	Id       *string `json:"id"  validate:"required"`
}
