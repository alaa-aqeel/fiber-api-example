package models 


import (
	"github.com/jinzhu/gorm"
)


type PermissionModel struct{
	gorm.Model 

	Name  string `sql:"not null;unique" json:"name"`
	Roles []RoleModel `gorm:"many2many:roles_permissions;"`
}


func (model *PermissionModel) All() (perms []PermissionModel) {

	Sql.Find(&perms)
	return 
}

func (model *PermissionModel) Get(id interface{}) (perm PermissionModel, err error) {

	err = Sql.First(&perm, id).Error
	return 
}

func (model *PermissionModel) GetRole(id interface{}) (role RoleModel, err error){

	err = Sql.Model(&model).Where("id = ?",id).Related(&role, "Roles").Error 
	return 
}

func (model *PermissionModel) AppendRole(role RoleModel) error {

	return Sql.Model(&model).Association("Roles").Append(role).Error
}

func (model *PermissionModel) DeleteRole(role RoleModel) error {

	return Sql.Model(&model).Association("Roles").Delete(role).Error
}
