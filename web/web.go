package web

import (
	"docker-file-management/internal/config"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Routes(e *echo.Echo) {

	t := &Template{
		templates: template.Must(template.ParseGlob("www/*.html")),
	}

	e.Renderer = t

	// Create a group for /
	g := e.Group("/")
	g.GET("", userPanel)
	g.GET("files/:file", userFile)

	// Create a group for /admin
	g = e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "admin" && password == config.GetConfig().AdminPassword {
			return true, nil
		}
		return false, nil
	}))
	g.GET("", adminPanel)
}
