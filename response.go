package inertia

import (
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, component string, props map[string]interface{}) error {
	if props == nil {
		props = fiber.Map{}
	}
	page := &Page{
		Component: component,
		Props:     props,
		URL:       c.Path(),
		Version:   version(),
	}

	if c.Get("X-Inertia") != "" {
		c.Set("X-Inertia", "true")

		return c.JSON(page)
	}

	return c.Render("app", fiber.Map{
		"reactRefresh": reactRefresh(),
		"scripts":      vite([]string{"resources/js/app.tsx", "resources/js/Pages/" + page.Component + ".tsx"}),
		"inertia":      inertiaBody(page),
	})
}
