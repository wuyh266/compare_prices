package service

import (
	"compare_prices/server_go/dao"
	"compare_prices/server_go/model"
	"github.com/google/uuid"
)

type RecognizeService struct {
	taskDAO *dao.RecognizeTaskDAO
}

func NewRecognizeService(taskDAO *dao.RecognizeTaskDAO) *RecognizeService {
	return &RecognizeService{taskDAO: taskDAO}
}

func (s *RecognizeService) CreateRecognizeTask(category string, attributes map[string]interface{}, rawImageURL string) (*model.RecognizeResult, error) {
	taskID := "task_" + uuid.New().String()

	task := &model.RecognizeTask{
		TaskID:      taskID,
		Category:    category,
		Attributes:  attributes,
		RawImageURL: rawImageURL,
	}

	err := s.taskDAO.Create(task)
	if err != nil {
		return nil, err
	}

	suggestCards := []model.SuggestCard{
		{ID: "card_lowest_price", Title: "查看同款低价", Icon: "💰"},
		{ID: "card_official_store", Title: "只看旗舰店", Icon: "🏪"},
		{ID: "card_hot_recommend", Title: "相似爆款推荐", Icon: "🔥"},
	}

	return &model.RecognizeResult{
		TaskID:       taskID,
		Category:     category,
		Attributes:   attributes,
		SuggestCards: suggestCards,
	}, nil
}
