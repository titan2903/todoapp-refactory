package main

import (
	"fmt"
	"log"
	"os"
	"todoapp-refactory/auth"
	"todoapp-refactory/handler"
	"todoapp-refactory/middleware"
	"todoapp-refactory/todo"
	"todoapp-refactory/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsnMaster := os.Getenv("DATABASE_URL")
	db, errMaster := gorm.Open(postgres.Open(dsnMaster), &gorm.Config{})
	if errMaster != nil {
		log.Panic(errMaster)
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&user.User{}, &todo.Todo{})

	//! Auth
	authService := auth.NewService()

	//! Users
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	//! Todos
	todoRepository := todo.NewRepository(db)
	todoService := todo.NewService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	router := gin.Default()
	router.Use(middleware.CORSMiddleware()) // ! Allow cors

	api := router.Group("/api/v1")

	//! Router Users
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailHasBeenRegister)                                  //! Reouter check email sudah terdaftar di database atau belum (sudah = false, belum = true)
	api.GET("/users/fetch", middleware.AuthMiddleware(authService, userService), userHandler.FetchUser) //! Router check user yang login

	// ! Router Todos
	api.POST("/todos", middleware.AuthMiddleware(authService, userService), todoHandler.CreateTodo)
	api.GET("/todos", middleware.AuthMiddleware(authService, userService), todoHandler.GetTodos)
	api.GET("/todos/:id", middleware.AuthMiddleware(authService, userService), todoHandler.GetTodo)
	api.PUT("/todos/:id", middleware.AuthMiddleware(authService, userService), todoHandler.UpdateTodo) //! update Todo jika todo sudah selesai maka is_completed = 1 (0: belum selesai, 1: sudah selesai)
	api.DELETE("/todos/:id", middleware.AuthMiddleware(authService, userService), todoHandler.DeleteTodo)

	var port string
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "5000"
	}

	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
	// router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
