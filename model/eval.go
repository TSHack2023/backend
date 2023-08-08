package model

type Eval struct {
	EvalId      uint   `json:"eval_id" gorm:"primaryKey"`
	FileId      uint   `json:"file_id" gorm:"not null"`
	Evalname    string `json:"evalname" gorm:"not null"`
	Evalmin     uint   `json:"evalmin" gorm:"not null"`
	Evalmax     uint   `json:"evalmax" gorm:"not null"`
	Explanation string `json:"explanation"`
}

type EvalResponse struct {
	EvalId      uint   `json:"eval_id" gorm:"primaryKey"`
	Evalname    string `json:"evalname" gorm:"not null"`
	Evalmin     uint   `json:"evalmin" gorm:"not null"`
	Evalmax     uint   `json:"evalmax" gorm:"not null"`
	Explanation string `json:"explanation"`
}
