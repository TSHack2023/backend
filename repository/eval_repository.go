package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type IEvalRepository interface {
	GetAllEvals(evals *[]model.Eval, fileId uint) error
	GetEvalById(eval *model.Eval, evalId uint) error
	CreateEval(eval *model.Eval) error
	// UpdateTask(task *model.Task, userId uint, taskId uint) error
	// DeleteTask(userId uint, taskId uint) error
}

type evalRepository struct {
	db *gorm.DB
}

func NewEvalRepository(db *gorm.DB) IEvalRepository {
	return &evalRepository{db}
}

func (er *evalRepository) GetAllEvals(evals *[]model.Eval, fileId uint) error {
	if err := er.db.Where("file_id=?", fileId).Order("eval_id").Find(evals).Error; err != nil {
		return err
	}
	return nil
}

func (er *evalRepository) GetEvalById(eval *model.Eval, evalId uint) error {
	if err := er.db.Where("eval_id=?", evalId).First(eval).Error; err != nil {
		return err
	}
	return nil
}

func (er *evalRepository) CreateEval(eval *model.Eval) error {
	if err := er.db.Create(eval).Error; err != nil {
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
