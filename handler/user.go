package handler

import (
	"log"
	"new-bank-api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

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
	return c.JSON(results)
}
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &model.User{}
	mgm.Coll(user).FindByID(id, user)

	return c.JSON(user)
}
func CreateUser(c *fiber.Ctx) error {
	user := new(model.User)
	c.BodyParser(user)
	if err := mgm.Coll(user).Create(user); err != nil {
		return c.JSON(err)
	}

	return c.JSON(fiber.Map{"status": "create user successed!"})
}
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &model.User{}
	mgm.Coll(user).FindByID(id, user)

	newData := new(model.User)
	c.BodyParser(newData)

	user.Cpf = newData.Cpf
	user.Debit = newData.Debit

	if err := mgm.Coll(user).Update(user); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "Update user successed!", "user": user})
}
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &model.User{}
	mgm.Coll(user).FindByID(id, user)

	if err := mgm.Coll(user).Delete(user); err != nil {
		return err
	}
	return c.JSON(fiber.Map{"status": "delete user successed!", "user": user})
}
