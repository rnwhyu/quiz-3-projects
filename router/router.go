package routers

import (
	"project-quiz3/auth"
	"project-quiz3/controllers"
	"os"
	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

func Setup(router *gin.Engine) {
	//bangun datar
	router.GET("persegi", controllers.HitungPersegi)
	router.GET("persegi-panjang", controllers.HitungPP)
	router.GET("segitiga-sama-sisi", controllers.HitungSegitiga)
	router.GET("lingkaran", controllers.HitungLingkaran)
	//book
	router.GET("/books", controllers.FindAllBooks)
	router.POST("/books", auth.BasicAuth(), controllers.CreateBook)
	router.PUT("/books/:id", auth.BasicAuth(), controllers.UpdateBook)
	router.DELETE("/books/:id", auth.BasicAuth(), controllers.DeleteBook)
	//category
	router.GET("/category", controllers.FindAll)
	router.GET("/category/:id", controllers.FindByID)
	router.GET("/category/:id/books", controllers.FindBooks)
	router.POST("/category", auth.BasicAuth(), controllers.CreateCategory)
	router.PUT("/category:id", auth.BasicAuth(), controllers.UpdateCategory)
	router.DELETE("/category/:id", auth.BasicAuth(), controllers.DeleteCategory)
	router.Run(":"+os.Getenv("PORT"))
}
