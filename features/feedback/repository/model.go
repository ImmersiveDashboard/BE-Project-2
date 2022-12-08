package repository

import (
	_feedback "immersive-dashboard/features/feedback"

	"gorm.io/gorm"
)

// struct gorm model

type Feedback struct {
	gorm.Model
	MenteeID       uint
	UserID         uint
	Description    string
	Mentee_Status  string
	Changed_Status string
	Mentees        Mentee
	Users          User
}

type User struct {
	gorm.Model
	Full_Name  string
	Email      string
	Password   string
	TeamID     uint // <<tidak ada underscore (default setting)
	Role       string
	Status     string
	Permission string
	Feedbacks  []Feedback
}

type Mentee struct {
	gorm.Model
	Mentee_Name           string
	ClassID               uint
	Status                string
	Category              string
	Gender                string
	Current_Address       string
	Home_Address          string
	Email                 string
	Telegram              string
	Phone                 string
	Emergency_Name        string
	Emergency_Phone       string
	Emergency_Status      string
	Education_Type        string
	Education_Major       string
	Education_Graduate    string
	Education_Institution string
	Feedbacks             []Feedback
}

// mengubah struct core ke struct model gorm
func fromCore(dataCore _feedback.CoreFeedback) Feedback {
	feedbackGorm := Feedback{
		MenteeID:       dataCore.MenteeID,
		UserID:         dataCore.UserID,
		Description:    dataCore.Description,
		Mentee_Status:  dataCore.Mentee_Status,
		Changed_Status: dataCore.Changed_Status,
	}
	return feedbackGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Feedback) toCore() _feedback.CoreFeedback {
	return _feedback.CoreFeedback{
		ID: dataModel.ID,
		Mentees: _feedback.CoreMentee{
			ID:          dataModel.MenteeID,
			Mentee_Name: dataModel.Mentees.Mentee_Name,
		},
		Users: _feedback.CoreUser{
			ID:        dataModel.UserID,
			Full_Name: dataModel.Users.Full_Name,
		},
		Description:    dataModel.Description,
		Mentee_Status:  dataModel.Mentee_Status,
		Changed_Status: dataModel.Changed_Status,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Feedback) []_feedback.CoreFeedback {
	var dataCore []_feedback.CoreFeedback
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}