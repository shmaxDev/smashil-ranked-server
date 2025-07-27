package dtos

type ReportMatchDto struct {
    Player1Id    *string `json:"player1Id" validate:"required"`
    Player2Id    *string `json:"player2Id" validate:"required"`
    Player1Score *int    `json:"player1Score" validate:"required"`
    Player2Score *int    `json:"player2Score" validate:"required"`
}