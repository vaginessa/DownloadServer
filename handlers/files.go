package handlers

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/dls/entities"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func FilesHandler(c *fiber.Ctx) error {
	content, err := ioutil.ReadDir("./files")
	if err != nil {
		log.Fatal(err)
	}

	var files []entities.File
	for _, file := range content {
		if file.IsDir() {
			continue
		}
		files = append(files, entities.File{
			Name: file.Name(),
			Size: utils.ByteSize(uint64(file.Size())),
			ModTime: file.ModTime().Format(time.RFC822),
		})
	}

	return c.Render("index", entities.Content{
		Files: files,
	})
}