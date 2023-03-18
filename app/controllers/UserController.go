package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	return c.SendString("GetAllUsers")
}

func (uc *UserController) GetUser(c *fiber.Ctx) error {
	return c.SendString("GetUser")
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	return c.SendString("CreateUser")
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	return c.SendString("UpdateUser")
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	return c.SendString("DeleteUser")
}
