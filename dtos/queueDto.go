package dtos

type Queue struct {
	UserId *string `json:"userId"`
	Elo    *int    `json:"elo"`
}
