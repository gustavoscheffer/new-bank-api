package handler

import (
	"log"
	"new-bank-api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func findUserCpf(cpf int64) *model.User {
	user := new(model.User)
	mgm.Coll(user).First(bson.M{"cpf": cpf}, user)
	return user
}
func GetAllUser(c *fiber.Ctx) error {

	coll := mgm.Coll(&model.User{})
	cursor, err := coll.Find(mgm.Ctx(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var results []model.User

	if err := cursor.All(mgm.Ctx(), &results); err != nil {
		log.Fatal(err)
	}
	if results == nil {
		c.JSON(fiber.Map{"status": "There's no user to list"})
		return c.SendStatus(200)
	}
	return c.JSON(results)
}
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &model.User{}

	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		c.JSON(fiber.Map{"status": "User not found"})
		return c.SendStatus(404)
	}
	return c.JSON(fiber.Map{"status": "User found", "result": user})
}
func CreateUser(c *fiber.Ctx) error {

	user := &model.User{}
	c.BodyParser(user)

	if user.Cpf == findUserCpf(user.Cpf).Cpf {
		c.JSON(fiber.Map{"status": "user user already exists"})
		return c.SendStatus(400)
	}

	if err := mgm.Coll(user).Create(user); err != nil {
		c.JSON(fiber.Map{"status": "user cannot be created"})
		return c.SendStatus(500)
	}

	c.JSON(fiber.Map{"status": "user created!", "result": user})
	return c.SendStatus(201)
}
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &model.User{}

	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		c.JSON(fiber.Map{"status": "User not found"})
		return c.SendStatus(404)
	}

	c.BodyParser(user)

	if err := mgm.Coll(user).Update(user); err != nil {
		c.JSON(fiber.Map{"status": "User cannot be updated"})
		return c.SendStatus(500)
	}

	return c.JSON(fiber.Map{"status": "User up to date", "user": user})
}
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &model.User{}

	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		c.JSON(fiber.Map{"status": "User not found"})
		return c.SendStatus(404)
	}
	err := mgm.Coll(user).Delete(user)
	if err != nil {
		c.JSON(fiber.Map{"status": "User cannot be deleted"})
		return c.SendStatus(500)
	}
	return c.JSON(fiber.Map{"status": "User deleted", "user": user})
}
