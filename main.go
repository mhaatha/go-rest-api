package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/mhaatha/go-rest-api/app"
	"github.com/mhaatha/go-rest-api/controller"
	"github.com/mhaatha/go-rest-api/repository"
	"github.com/mhaatha/go-rest-api/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbPassword := os.Getenv("DB_PASSWORD")

	db := app.NewDB(dbPassword)
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
