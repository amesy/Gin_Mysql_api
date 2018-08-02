package main

import (
	db "gin_exercise/database"
	"gin_exercise/routers"
)

func main() {
	defer db.SqlDB.Close()
	r := routers.InitRouter()
	r.Run(":9999")
}
