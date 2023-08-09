package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type IScoreRepository interface {
	GetAllScores(scores *[]model.Score, evalId uint) error
	GetScoreById(score *model.Score, scoreId uint) error
	CreateScore(score *model.Score) error
	GetScoreByUsername(scores *[]model.Score, username string) error
	GetUsersByEvalId(users *[]string, evalId uint) error
}

type scoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) IScoreRepository {
	return &scoreRepository{db}
}

func (sr *scoreRepository) GetAllScores(scores *[]model.Score, evalId uint) error {
	if err := sr.db.Where("eval_id=?", evalId).Order("score_id").Find(scores).Error; err != nil {
		return err
	}
	return nil
}

func (sr *scoreRepository) GetScoreById(score *model.Score, scoreId uint) error {
	if err := sr.db.Where("score_id=?", scoreId).First(score).Error; err != nil {
		return err
	}
	return nil
}

func (sr *scoreRepository) CreateScore(score *model.Score) error {
	if err := sr.db.Create(score).Error; err != nil {
		return err
	}
	return nil
}

func (sr *scoreRepository) GetScoreByUsername(scores *[]model.Score, username string) error {
	if err := sr.db.Where("username = ?", username).Find(scores).Error; err != nil {
		return err
	}
	return nil
}

func (sr *scoreRepository) GetUsersByEvalId(users *[]string, evalId uint) error {
	if err := sr.db.Model(&model.Score{}).Where("eval_id = ?", evalId).Pluck("username", users).Error; err != nil {
		return err
	}
	return nil
}

/*
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
*/
