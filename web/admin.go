package web

import (
	"docker-file-management/internal/config"

	"github.com/labstack/echo/v4"
)

// Admin route with Basic Auth from env
// Path: web/admin.go
func adminPanel(c echo.Context) error {
	// Check basic auth
	username, password, ok := c.Request().BasicAuth()
	if !ok || username != "admin" || password != config.GetConfig().AdminPassword {
		return c.JSON(403, map[string]interface{}{
			"message": "forbidden",
		})
	}

	// return www/admin.html
	return c.File("www/admin.html")
}
