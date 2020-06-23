package main 

import (
	"flag"
	"github.com/alaaProg/postapi/app"
	"github.com/alaaProg/postapi/models"
)

func main(){

	serverApp := app.CreateApp() // Create Application fiber 

	db := app.InitDatabase("db.sqlite") // Setup Connect Database `sqlite3 `
	defer db.Close()
	
	run     := flag.Bool("run", false, "run fiber server host ")
	host    := flag.String("host", "127.0.0.1:8080", "listen default 127.0.0.1:8080")
	migrate := flag.Bool("migrate", false, "Migrate Database")
	reset   := flag.Bool("reset", false, "Reset Database")
	seed    := flag.Bool("seed", false, "Insert default data")


	flag.Parse() 

	if *migrate { 

		models.CreateTables() // Create All Tables 

	} else if  *reset {

		models.DropTables() // Delete All Table 
		models.CreateTables()  

	}else if *seed {

		models.SeedeTable()

	} else if  *run { 

		serverApp.Listen(*host) // Run Server Listen 
	}

}


