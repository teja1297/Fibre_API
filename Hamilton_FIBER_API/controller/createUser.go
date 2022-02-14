package controller

import (
	"HAMILTON_FIBER_API/database"
	"HAMILTON_FIBER_API/entity"
	"context"
	"encoding/json"

	"github.com/gofiber/fiber"
)

func CreateUser(c *fiber.Ctx) {
	collection, err := database.GetMongoDbCollection("TEST1", "User")
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var user entity.User
	json.Unmarshal([]byte(c.Body()), &user)

	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}

