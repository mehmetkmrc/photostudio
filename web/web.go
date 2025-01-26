package web

import(
	"github.com/gofiber/fiber/v3"
)

func IndexPage(c fiber.Ctx) error {
	path := "home"
	return c.Render(path, fiber.Map{
		"Title": "Ismail Kömürcü",
	},"layout/main")
}

func ContactPage(c fiber.Ctx) error {
	path := "contact"
	return c.Render(path, fiber.Map{
		"Title": "İletişim",
	}, "layout/main")
}

func AboutUsPage(c fiber.Ctx) error {
	path := "about-us"
	return c.Render(path, fiber.Map{
		"Title": "Hakkında",
	}, "layout/main")
}

func PhotographyPage(c fiber.Ctx) error {
	path := "photography"
	return c.Render(path, fiber.Map{
		"Title": "Koleksiyonlar",
	}, "layout/main")
}


