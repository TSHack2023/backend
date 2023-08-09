package model

type Score struct {
	ScoreId  uint   `json:"score_id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null"`
	EvalId   uint   `json:"eval_id" gorm:"not null"`
	Score    uint   `json:"score" gorm:"not null"`
}

type ScoreResponse struct {
	ScoreId  uint   `json:"score_id" gorm:"primaryKey"`
	Evalname string `json:"evalname" gorm:"not null"`
	Score    uint   `json:"score" gorm:"not null"`
}

type AnswerResponse struct {
	Username          string          `json:"username" gorm:"primaryKey"`
	ScoreResponseList []ScoreResponse `json:"scorelist"`
}

type FinalAnswerResponse struct {
	Result bool             `json:"result"`
	Answer []AnswerResponse `json:"answerlist"`
}
