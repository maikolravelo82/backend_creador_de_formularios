package modelos

import (
	"errors"

	"github.com/Grimmore/Server/database"
	"github.com/jinzhu/gorm"
)

type Area struct {
	gorm.Model
	Nombre   string
	Valor    int
	Answerid string `sql:"index"`
}

func GetAllArea() []Area {
	var area []Area
	db := database.GetDB()
	db.Find(&area)
	return area
}

// Crea un objeto evento y lo añade como tupla a la base da datos
func NewArea(n string, v int, i string) (*Area, error) {
	if n == "" || v == 0 || i == "" {
		return nil, errors.New("empty fields")
	}

	Area := Area{
		Nombre:   n,
		Valor:    v,
		Answerid: i,
	}

	db := database.GetDB()
	err := db.Create(&Area).Error
	return &Area, err
}

// borra una instancia de objeto evento con la id pasada por parametro
func DeleteArea(id string) error {
	db := database.GetDB()
	return db.Where("id == ?", id).Delete(&Area{}).Error
}

func GetAreaByPath(p string) (*Area, error) {
	db := database.GetDB()
	var e Area
	err := db.Model(&e).Where("id == ?", p).First(&e).Error
	return &e, err
}

func GetAreaByAnswer(p string) (*[]Area, error) {
	db := database.GetDB()
	var e []Area
	err := db.Model(&e).Where("answerid == ?", p).Find(&e).Error
	return &e, err
}
