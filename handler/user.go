package handler

import (
	"fmt"
	"log"
	"net/http"

	"new-bank-api/model"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func findUserCpf(cpf string) *model.User {
	user := &model.User{}
	mgm.Coll(user).First(bson.M{"cpf": cpf}, user)
	return user
}

// GetAllUser is a function to get all users data from database
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]model.User}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/user [get]
func GetAllUser(c *fiber.Ctx) error {

	coll := mgm.Coll(&model.User{})
	cursor, err := coll.Find(mgm.Ctx(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var results []model.User

	if err := cursor.All(mgm.Ctx(), &results); err != nil {

		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get all users",
		Data:    results,
	})
}

// GetUser is a function to get a user by ID
// @Summary Get user by ID
// @Description Get user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} ResponseHTTP{data=[]model.User}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/user/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &model.User{}

	if err := mgm.Coll(user).FindByID(id, user); err != nil {

		switch err.Error() {
		case "the provided hex string is not a valid ObjectID":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("User with id %v not found.", id),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}

	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get user by ID.",
		Data:    user,
	})
}

// CreateUser create a new user data
// @Summary Create a new user
// @Description Create user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.User true "Create user"
// @Success 200 {object} ResponseHTTP{data=model.User}
// @Failure 400 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/user [post]
func CreateUser(c *fiber.Ctx) error {

	user := &model.User{}
	c.BodyParser(user)

	if user.Cpf == findUserCpf(user.Cpf).Cpf {
		c.JSON(fiber.Map{"status": "user user already exists"})
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: fmt.Sprintf("User with cpf: %v already exists.", user.Cpf),
			Data:    nil,
		})
	}

	if err := mgm.Coll(user).Create(user); err != nil {
		c.JSON(fiber.Map{"status": "user cannot be created"})
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(http.StatusCreated).JSON(ResponseHTTP{
		Success: true,
		Message: "Success create a user",
		Data:    user,
	})
}

// UpdateUser create a new user data
// @Summary Update a user
// @Description Update user
// @Tags user
// @Accept json
// @Produce json
// @Param user body model.User true "Update user"
// @Param id path string true "user ID"
// @Success 200 {object} ResponseHTTP{data=model.User}
// @Failure 400 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/user/{id} [patch]
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &model.User{}

	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		switch err.Error() {
		case "the provided hex string is not a valid ObjectID":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("User with id %v not found.", id),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}
	}

	c.BodyParser(user)

	if err := mgm.Coll(user).Update(user); err != nil {
		c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User with id: %v up-to-date", id),
		Data:    user,
	})
}

// DeleteUser function removes a user by ID
// @Summary Remove user by ID
// @Description Remove user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} ResponseHTTP{}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/user/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user := &model.User{}

	if err := mgm.Coll(user).FindByID(id, user); err != nil {
		switch err.Error() {
		case "the provided hex string is not a valid ObjectID":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("User with id %v not found.", id),
				Data:    nil,
			})
		case "mongo: no documents in result":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("User with id %v not found.", id),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}
	}
	err := mgm.Coll(user).Delete(user)
	if err != nil {

		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: fmt.Sprintf("User with id: %v deleted", id),
		Data:    user,
	})
}
