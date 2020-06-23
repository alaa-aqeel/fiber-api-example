package models 


import (

	"github.com/jinzhu/gorm"
)


type RoleModel struct{
	gorm.Model 

	Name        string            `sql:"not null;unique" json:"name" required:"true"`
	Users       []UserModel       `gorm:"ForeignKey:RoleId" json:"users"`
	Permissions []*PermissionModel `gorm:"many2many:roles_permissions;"`
}

func (model *RoleModel) HashPermission(name string) bool {

	var perm PermissionModel
	err := Sql.Model(&model).Where("name = ?", name).Related(&perm, "Permissions").Error 
	
	return err == nil 
}

func (model *RoleModel) GetUser(id interface{}) (users UserModel, err error) {

	err = Sql.Model(&model).Where("id = ?",id).Related(&users, "Users").Error 
	return
}	

func (model *RoleModel) DeleteUser(user UserModel) error {

	return Sql.Model(&model).Association("Users").Delete(user).Error
}


func (model *RoleModel) AppendUser(user UserModel) error {

	return Sql.Model(&model).Association("Users").Append(user).Error 
}	

func (model *RoleModel) Get(id interface{}) (role RoleModel, err error) {


	err = Sql.Preload("Permissions").Preload("Users").First(&role, id).Error
	return 
}

func (model *RoleModel) All() (roles []RoleModel) {

	Sql.Preload("Users").Find(&roles)
	return 
}

func (model *RoleModel) Create() error {

	return Sql.Create(&model).Error
}

func (model *RoleModel) Save() error {

	return Sql.Save(&model).Error
}

func (model *RoleModel) Delete() error{

	return Sql.Unscoped().Delete(&model).Error
}

