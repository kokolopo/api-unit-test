package main

import (
	"api_unit_test/controller"
	"api_unit_test/user"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// koneksi database
	dsn := "root:@tcp(127.0.0.1:3306)/unit_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("koneksi database gagal!!!")
	}

	// depedency layer user
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userController := controller.NewUserHandler(userService)

	// end point
	e := echo.New()

	userApi := e.Group("/api/v1")

	userApi.GET("/users", userController.GetAll)
	userApi.GET("/users/:id", userController.GetById)
	userApi.POST("/users", userController.RegisterUser)
	userApi.PUT("/users/:id", userController.UpdateUser)
	userApi.DELETE("/users/:id", userController.DeleteUser)

	e.Logger.Fatal(e.Start(":8888"))
}
