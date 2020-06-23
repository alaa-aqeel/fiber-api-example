package models

import "github.com/jinzhu/gorm"

var Sql *gorm.DB 

var CurrentUser UserModel


func CreateTables(){

	Sql.CreateTable(&UserModel{})
	Sql.CreateTable(&RoleModel{})
	Sql.CreateTable(&PermissionModel{})
}

func DropTables(){

	Sql.DropTable(&UserModel{})
	Sql.DropTable(&RoleModel{})
	Sql.DropTable(&PermissionModel{})
}


func SeedeTable(){

	admin := UserModel{
		Username:"admin",
		Password:"12345678",
	}

	Sql.Create(admin)


	Sql.Create(&RoleModel{
			Name:"admin",
			Users:[]UserModel{
				admin,
			},
		},
	)
	
	Sql.Create(&RoleModel{Name:"user",})

	Sql.Create(&PermissionModel{Name:"read",})
	Sql.Create(&PermissionModel{Name:"write",})
	Sql.Create(&PermissionModel{Name:"delete",})
	Sql.Create(&PermissionModel{Name:"update",})
}


