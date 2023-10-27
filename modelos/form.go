package modelos

import (
	"errors"

	"github.com/Grimmore/Server/database"
	"github.com/jinzhu/gorm"
)

type Form struct {
	gorm.Model
	Nombre string
}

func GetAllForm() []Form {
	var form []Form
	db := database.GetDB()
	db.Find(&form)
	return form
}

func NewForm(n string) (*Form, error) {
	if n == "" {
		return nil, errors.New("empty fields")
	}

	Form := Form{
		Nombre: n,
	}

	db := database.GetDB()
	err := db.Create(&Form).Error
	return &Form, err
}

// borra una instancia de objeto evento con la id pasada por parametro
func DeleteForm(id string) error {
	db := database.GetDB()
	return db.Where("id == ?", id).Delete(&Form{}).Error
}

func GetFormByPath(p string) (*Form, error) {
	db := database.GetDB()
	var e Form
	err := db.Model(&e).Where("id == ?", p).First(&e).Error
	return &e, err
}
