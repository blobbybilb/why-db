package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"whydb/config"
	"whydb/filestore"
	"whydb/postgresstore"
	"whydb/types"
)

func getStore() types.Store {
	switch config.StoreType {
	case "file":
		return filestore.NewFileStore()
	case "postgres":
		return postgresstore.NewPostgresStore()
	default:
		panic("invalid StoreType in config, please set a valid StoreType")
	}
}

func main() {
	app := fiber.New()
	store := getStore()

	app.Get("/set/:cat/:key/*", func(c *fiber.Ctx) error {
		cat, key, data := c.Params("cat"), c.Params("key"), c.Params("*")
		err := store.Set(cat, key, data)
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString("done")
	})

	app.Post("/set/:cat/:key/", func(c *fiber.Ctx) error {
		cat, key := c.Params("cat"), c.Params("key")
		err := store.Set(cat, key, string(c.Body()))
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString("done")
	})

	app.Get("/get/:cat/:key/", func(c *fiber.Ctx) error {
		cat, key := c.Params("cat"), c.Params("key")
		text, err := store.Get(cat, key)
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString(string(text))
	})

	app.Get("/add/:cat/:key/*", func(c *fiber.Ctx) error {
		cat, key, data := c.Params("cat"), c.Params("key"), c.Params("*")
		err := store.Add(cat, key, data)
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString("done")
	})

	app.Post("/add/:cat/:key/*", func(c *fiber.Ctx) error {
		cat, key := c.Params("cat"), c.Params("key")
		err := store.Add(cat, key, string(c.Body()))
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString("done")
	})

	app.Get("/del/:cat/:key/", func(c *fiber.Ctx) error {
		cat, key := c.Params("cat"), c.Params("key")
		err := store.Del(cat, key)
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.SendString("done")
	})

	app.Get("/exists", func(c *fiber.Ctx) error {
		return c.SendString("exists - why?db 0.2")
	})

	_ = app.Listen(config.Address + ":" + fmt.Sprint(config.Port))
}
