package main

import (
	"flag"
	"log"

	"github.com/dls/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)



func main() {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/assets", "./assets")
	app.Static("/files", "./files")

	app.Get("/", handlers.FilesHandler)

	addr := flag.String("address", ":3000", "address to host on")
	log.Fatal(app.Listen(*addr))
}