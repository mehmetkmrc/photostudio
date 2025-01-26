package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"time"
	"photostudio/web"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/gofiber/template/html/v2"
)

const (
	viewPath   = "./client/templates"
	publicPath = "./client/public"
	renderType = ".html"
)

func add(x, y int) int {
	return x + y
}

func main() {

	engine := html.New(viewPath, renderType)
	engine.AddFunc("unescape", func(s string) template.HTML {
		return template.HTML(s)
	})

	engine.AddFunc("safe", func(s string) template.HTML {
		return template.HTML(s) // HTML olarak işaretler, güvenli kabul eder
	})

	engine.AddFunc("attr", func(s string) template.HTMLAttr {
		return template.HTMLAttr(s) // Attribute olarak işaretler
	})
	engine.AddFunc("safeHTML", func(s string) template.HTML {
		return template.HTML(s) // HTML olarak işaretle
	})
	engine.AddFunc("raw", func(s string) template.HTML {
		return template.HTML(s) // Mark string as raw HTML
	})
	engine.AddFunc("add", add)
	app := fiber.New(fiber.Config{
		ReadTimeout:   time.Minute * time.Duration(5),
		StrictRouting: false,
		CaseSensitive: true,
		BodyLimit:     4 * 1024 * 1024,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		AppName:       "photostudio",
		Immutable:     true,
		Views:         engine,
		//ViewsLayout: "layouts/main",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			var e *fiber.Error
			if errors.As(err, &e) {
				if e.Code == fiber.StatusNotFound {
					return c.Render("404", fiber.Map{
						"Title": "Page Not Found",
					})
				}
				return c.Status(e.Code).Render("error", fiber.Map{
					"Title":   "Error",
					"Message": e.Message,
				})
			}
			return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{
				"Title":   "Internal Server Error",
				"Message": "An unexpected error occured.",
			})
		},
	})
	

	app.Use(static.New(publicPath))


	app.Get("/", web.IndexPage)
	app.Get("/contact", web.ContactPage)
	app.Get("/about-us", web.AboutUsPage)
	app.Get("/collections", web.PhotographyPage)


	log.Fatal(app.Listen("0.0.0.0:3030"))
}