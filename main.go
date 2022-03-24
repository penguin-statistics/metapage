package main

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/kelseyhightower/envconfig"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

func main() {
	var conf Config
	err := envconfig.Process("penguinmeta", &conf)
	if err != nil {
		panic(err)
	}

	engine := html.New("./web/dist", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	body := blackfriday.MarkdownBasic([]byte(conf.Index.Body))
	body = bluemonday.UGCPolicy().SanitizeBytes(body)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": conf.Index.Title,
			"Body":  template.HTML(string(body)),
		})
	})

	app.Listen("127.0.0.1:8080")
}
