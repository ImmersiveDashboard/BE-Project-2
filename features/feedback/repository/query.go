package repository

import (
	"errors"
	"immersive-dashboard/features/feedback"

	"gorm.io/gorm"
)

type feedbackRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.RepositoryInterface {
	return &feedbackRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *feedbackRepository) Create(input feedback.CoreFeedback) (row int, err error) {
	feedbackGorm := fromCore(input)
	tx := repo.db.Create(&feedbackGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements user.Repository
func (repo *feedbackRepository) GetAll() (data []feedback.CoreFeedback, err error) {
	var feedbacks []Feedback

	tx := repo.db.Find(&feedbacks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(feedbacks)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (*feedbackRepository) GetById(id int) (data feedback.CoreFeedback, err error) {
	panic("unimplemented")
}
