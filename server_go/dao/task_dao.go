package dao

import (
	"compare_prices/server_go/model"

	"gorm.io/gorm"
)

type RecognizeTaskDAO struct {
	db *gorm.DB
}

func NewRecognizeTaskDAO(db *gorm.DB) *RecognizeTaskDAO {
	return &RecognizeTaskDAO{db: db}
}

func (d *RecognizeTaskDAO) GetByTaskID(taskID string) (*model.RecognizeTask, error) {
	var task model.RecognizeTask
	err := d.db.Where("task_id = ?", taskID).First(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (d *RecognizeTaskDAO) Create(task *model.RecognizeTask) error {
	return d.db.Create(task).Error
}
