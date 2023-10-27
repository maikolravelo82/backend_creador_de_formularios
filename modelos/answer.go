package modelos

import (
	"errors"

	"github.com/Grimmore/Server/database"
	"github.com/jinzhu/gorm"
)

type Answer struct {
	gorm.Model
	Context    string
	Questionid string
}

func GetAllAnswer() []Answer {
	var answer []Answer
	db := database.GetDB()
	db.Find(&answer)
	return answer
}

func NewAnswer(n, i string) (*Answer, error) {
	if n == "" || i == "" {
		return nil, errors.New("empty fields")
	}

	Answer := Answer{
		Context:    n,
		Questionid: i,
	}

	db := database.GetDB()
	err := db.Create(&Answer).Error
	return &Answer, err
}

// borra una instancia de objeto evento con la id pasada por parametro
func DeleteAnswer(id string) error {
	db := database.GetDB()
	return db.Where("id == ?", id).Delete(&Answer{}).Error
}

func GetAnswerByPath(p string) (*Answer, error) {
	db := database.GetDB()
	var e Answer
	err := db.Model(&e).Where("id == ?", p).First(&e).Error
	return &e, err
}

func GetAnswerByQuestion(p string) (*[]Question, error) {
	db := database.GetDB()
	var e []Question
	err := db.Model(&e).Where("questionid == ?", p).Find(&e).Error
	return &e, err
}
