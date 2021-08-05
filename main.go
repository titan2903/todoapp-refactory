package main

import (
	"OAuth/auth"
	"OAuth/handler"
	"OAuth/middleware"
	"OAuth/todo"
	"OAuth/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main()  {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPassword := myEnv["DB_PASSWORD"]
	dbHost := myEnv["DB_HOST"]
	dbName := myEnv["DB_NAME"]
	dsn := fmt.Sprintf("root:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbPassword, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

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
	api.POST("/email_checkers", userHandler.CheckEmailHasBeenRegister) //! Reouter check email sudah terdaftar di database atau belum (sudah = false, belum = true)
	api.GET("/users/fetch", middleware.AuthMiddleware(authService, userService), userHandler.FetchUser) //! Router check user yang login

	// ! Router Todos
	api.POST("/todos", middleware.AuthMiddleware(authService, userService), todoHandler.CreateTodo)
	api.GET("/todos", middleware.AuthMiddleware(authService, userService), todoHandler.GetTodos)
	api.GET("/todos/:id", middleware.AuthMiddleware(authService, userService), todoHandler.GetTodo)
	api.PUT("/todos/:id", middleware.AuthMiddleware(authService, userService), todoHandler.UpdateTodo) //! update Todo jika todo sudah selesai maka is_completed = 1 (0: belum selesai, 1: sudah selesai)
	api.DELETE("/todos/:id", middleware.AuthMiddleware(authService, userService), todoHandler.DeleteTodo)
	
	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}