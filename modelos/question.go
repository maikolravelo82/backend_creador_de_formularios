package modelos

import (
	"errors"

	"github.com/Grimmore/Server/database"
	"github.com/jinzhu/gorm"
)

type Question struct {
	gorm.Model
	Context string
	Formid  string
}

func GetAllQuestion() []Question {
	var question []Question
	db := database.GetDB()
	db.Find(&question)
	return question
}

func NewQuestion(n, i string) (*Question, error) {
	if n == "" {
		return nil, errors.New("empty fields")
	}

	Question := Question{
		Context: n,
		Formid:  i,
	}

	db := database.GetDB()
	err := db.Create(&Question).Error
	return &Question, err
}

// borra una instancia de objeto evento con la id pasada por parametro
func DeleteQuestion(id string) error {
	db := database.GetDB()
	return db.Where("id == ?", id).Delete(&Question{}).Error
}

func GetQuestionByForm(p string) (*[]Question, error) {
	db := database.GetDB()
	var e []Question
	err := db.Model(&e).Where("formid == ?", p).Find(&e).Error
	return &e, err
}

func GetQuestionByPath(p string) (*Question, error) {
	db := database.GetDB()
	var e Question
	err := db.Model(&e).Where("id == ?", p).First(&e).Error
	return &e, err
}
