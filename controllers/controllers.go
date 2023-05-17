package controllers

import (
	"encoding/json"

	"github.com/GiorgosMarga/go-fiber/database"
	"github.com/GiorgosMarga/go-fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	res, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	return c.Send(res)
}

func CreateUser(c *fiber.Ctx) error {
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	// set new uuid
	user.ID = uuid.New()
	database.DB.Create(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}
	database.DB.First(&user, "id=?", id)
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.DB.Save(&user)
	return c.JSON(user)
}
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := models.User{}
	database.DB.Delete(&user, "id= ?", id)
	return c.JSON(user)
}
func GetUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")
	database.DB.First(&user, "id= ?", id)
	return c.JSON(user)
}
