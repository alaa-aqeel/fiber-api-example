package app 


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/alaaProg/postapi/models"
)


var errorsql error 

func InitDatabase(dbname string) *gorm.DB {

	models.Sql, errorsql = gorm.Open("sqlite3", dbname)

	if errorsql != nil {

		panic(errorsql)
	}	

	
	models.Sql.LogMode(true)

	return models.Sql
}