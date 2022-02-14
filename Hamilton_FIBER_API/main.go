package main

import (

	//"fmt"

	"HAMILTON_FIBER_API/controller"

	"github.com/gofiber/fiber"
)

const port = 8000

func main() {
	app := fiber.New()
	app.Get("/User/:id?", controller.GetUser)
	app.Post("/User", controller.CreateUser)
	app.Post("/User/Deposit", deposit)
	app.Post("/User/Withdrawl", withdrawl)
	app.Listen(port)
}
func deposit(c *fiber.Ctx) {

	controller.DepositOrWithdrawl(c, true)
}
func withdrawl(c *fiber.Ctx) {
	//depositOrWithdrawl(c, false)
}
