package controller

import (
	"HAMILTON_FIBER_API/database"
	"HAMILTON_FIBER_API/entity"
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func DepositOrWithdrawl(c *fiber.Ctx, track bool) {

	_, err := database.GetMongoDbCollection("TEST1", "User")
	fmt.Println("1")
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var user entity.User
	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	update := bson.M{
		"$set": person,
	}
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	

	
	response, _ := json.Marshal(res)
	c.Send(response)
}
