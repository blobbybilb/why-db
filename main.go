package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"whydb/config"
	"whydb/filestore"
	"whydb/types"
)

func getStore() types.Store {
	switch config.StoreType {
	case "file":

		return filestore.NewFileStore()
	default:
		return filestore.NewFileStore()
	}
}

func main() {
	app := fiber.New()

	store := getStore()

	app.Get("/set/:cat/:key/*", func(c *fiber.Ctx) error {
		cat, key, data := c.Params("cat"), c.Params("key"), c.Params("*")
		store.Set(cat, key, data)
		return c.SendString("done")
	})

	app.Post("/set/:cat/:key/", func(c *fiber.Ctx) error {
		cat, key := c.Params("cat"), c.Params("key")
		store.Set(cat, key, string(c.Body()))
		return c.SendString("done")
	})

	app.Get("/get/:cat/:key/", func(c *fiber.Ctx) error {
		cat, key := c.Params("cat"), c.Params("key")
		text, _ := store.Get(cat, key)
		return c.SendString(string(text))
	})

	app.Get("/add/:cat/:key/*", func(c *fiber.Ctx) error {
		cat, key, data := c.Params("cat"), c.Params("key"), c.Params("*")
		store.Add(cat, key, data)
		return c.SendString("done")
	})

	app.Post("/add/:cat/:key/*", func(c *fiber.Ctx) error {
		cat, key := c.Params("cat"), c.Params("key")
		store.Add(cat, key, string(c.Body()))
		return c.SendString("done")
	})

	app.Get("/del/:cat/:key/", func(c *fiber.Ctx) error {
		cat, key := c.Params("cat"), c.Params("key")
		store.Del(cat, key)
		return c.SendString("done")
	})
	_ = app.Listen(config.Address + ":" + fmt.Sprint(config.Port))
}
