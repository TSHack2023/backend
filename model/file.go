package model

import "time"

type File struct {
	FileId    uint      `json:"file_id" gorm:"primaryKey"`
	Filename  string    `json:"filename" gorm:"not null"`
	Fileurl   string    `json:"file_url" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	Username  string    `json:"username"`
}

type FileResponse struct {
	FileId    uint      `json:"id" gorm:"primaryKey"`
	Filename  string    `json:"title" gorm:"not null"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
} //一覧時に個数だけ送る

type FileReviewResponse struct {
	Result   bool           `json:"result"`
	Fileurl  string         `json:"file_url"`
	Evallist []EvalResponse `json:"evallist"`
} //評価ページに移動したときに送る
