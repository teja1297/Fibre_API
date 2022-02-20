package controller

import (
	"HAMILTON_FIBER_API/database"
	"context"
	"fmt"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func DepositOrWithdrawl(c *fiber.Ctx, track bool) {
	var balance int32
	col, err := database.GetMongoDbCollection("TEST1", "User")
	fmt.Println("1")
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	id := c.Params("ID")
	objID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	if err != nil {
		fmt.Println("ObjectIDFromHex ERROR", err)
	} else {
		fmt.Println("ObjectIDFromHex:", objID)
	}
	if track {
		balance += int32(objID[balance])
	}
	if !track {
		balance -= int32(objID[balance])
	}
	update := bson.M{"$set": bson.M{"balance": balance}}
	result, _ := col.UpdateOne(
		context.Background(),
		filter,
		update,
	)

	c.Send(result)
}
