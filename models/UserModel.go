package models


import (
	"fmt"
	"github.com/jinzhu/gorm"
)


type UserModel struct {
  gorm.Model

  Username string `sql:"not null;unique" json:"username" required:"true"`
  Password string `sql:"not null" json:"password" required:"true"`
  RoleId   uint
}


func (model *UserModel) Role() (role RoleModel,err error) {

	err = Sql.Model(&RoleModel{}).First(&role, model.RoleId).Error 
	return 
}

func (model *UserModel) Create() error {

	return Sql.Create(&model).Error
}

func (model *UserModel) All() (users []UserModel) {

	Sql.Find(&users)
	return users
}

func (model *UserModel) GetBy(field string, value interface{}) (user UserModel, err error) {


	err = Sql.Where(fmt.Sprintf("`%s` = ?",field), value).First(&user).Error
	return
}

func (model *UserModel) Get(id interface{}) (user UserModel, err error) {

	if err = Sql.First(&user, id).Error; err != nil {

		return user, err
	}

	return user, nil
}

func (model *UserModel) Delete() error{

	return Sql.Unscoped().Delete(&model).Error
}

func (model *UserModel) Save() error {
	
	// Sql.Model(&model).Updates(user)
	return Sql.Save(&model).Error 
}	