package main

import (
	"database/sql"
	"fmt"
	"os"
	"project-quiz3/database"
	"project-quiz3/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
var (
	DB *sql.DB
	err error
)
func main(){
	//db.Setup()
	err = godotenv.Load("config/.env")
	if err!=nil{
		fmt.Println("Failed load file .env")
	}else{
		fmt.Println("Success read file .env")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"))
	fmt.Println(psqlInfo)
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}else{
		fmt.Println("Successfully connected to DB")
	}
	db.DbMigrate(DB)
	defer DB.Close()
	app :=gin.Default()
	routers.Setup(app)

}