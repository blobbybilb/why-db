package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
)

const dataDir = "data/"

func createDirIfNotExists(path string) {
	path = dataDir + path
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		_ = os.Mkdir(path, 0755)
	}
}

func main() {
	app := fiber.New()

	app.Get("/set/:dir/:file/*", func(c *fiber.Ctx) error {
		dir, file, data := c.Params("dir"), c.Params("file"), c.Params("*")
		filePath := "./" + dataDir + dir + "/" + file

		createDirIfNotExists(dir)
		_ = os.WriteFile(filePath, []byte(data), 0644)

		return c.SendString("done")
	})

	app.Post("/set/:dir/:file/", func(c *fiber.Ctx) error {
		// save POST data to file
		dir, file := c.Params("dir"), c.Params("file")
		filePath := "./" + dataDir + dir + "/" + file

		createDirIfNotExists(dir)
		_ = os.WriteFile(filePath, c.Body(), 0644)
		return c.SendString("done")
	})

	app.Get("/get/:dir/:file/", func(c *fiber.Ctx) error {
		dir, file := c.Params("dir"), c.Params("file")
		filePath := "./" + dataDir + dir + "/" + file
		text, _ := os.ReadFile(filePath)
		return c.SendString(string(text))
	})

	app.Get("/add/:dir/:file/*", func(c *fiber.Ctx) error {
		dir, file, data := c.Params("dir"), c.Params("file"), c.Params("*")
		filePath := "./" + dataDir + dir + "/" + file

		createDirIfNotExists(dir)
		appendFile, _ := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		_, _ = appendFile.WriteString(data)
		_ = appendFile.Close()

		return c.SendString("done")
	})

	app.Post("/add/:dir/:file/*", func(c *fiber.Ctx) error {
		dir, file := c.Params("dir"), c.Params("file")
		filePath := "./" + dataDir + dir + "/" + file

		createDirIfNotExists(dir)
		appendFile, _ := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		_, _ = appendFile.WriteString(string(c.Body()))
		_ = appendFile.Close()

		return c.SendString("done")
	})

	app.Get("/del/:dir/:file/", func(c *fiber.Ctx) error {
		dir, file := c.Params("dir"), c.Params("file")
		filePath := "./" + dataDir + dir + "/" + file

		_ = os.Remove(filePath)
		return c.SendString("done")
	})
	_ = app.Listen("0.0.0.0:34578")
}
