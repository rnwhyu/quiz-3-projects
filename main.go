package main

import (
	"database/sql"
	"project-quiz3/database"
	"project-quiz3/router"
	"github.com/gin-gonic/gin"
)
var (
	DB *sql.DB
)
func main(){
	db.Setup()
	db.DbMigrate(DB)
	defer DB.Close()
	app :=gin.Default()
	routers.Setup(app)

}