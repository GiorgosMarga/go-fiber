package main

import (
	"github.com/GiorgosMarga/go-fiber/controllers"
	"github.com/GiorgosMarga/go-fiber/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}
	app := fiber.New()
	v1 := app.Group("/api/v1/auth")
	v1.Get("/:id", controllers.GetUser)
	v1.Get("/", controllers.GetAllUsers)
	v1.Post("/", controllers.CreateUser)
	v1.Patch("/:id", controllers.UpdateUser)
	v1.Delete("/:id", controllers.DeleteUser)

	app.Listen(":3000")
}

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func initApp() error {
	err := loadENV()
	if err != nil {
		return err
	}

	err = database.ConnectToDB()
	if err != nil {
		return err
	}
	return nil
}
