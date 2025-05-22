package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kshitij/taskManager/router"
)

func main() {
	var app *fiber.App = router.RouteHandler();

	log.Fatal(app.Listen(":3000"))
}