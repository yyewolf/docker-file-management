package web

import (
	"docker-file-management/internal/config"

	"github.com/labstack/echo/v4"
)

func userPanel(c echo.Context) error {
	// Check X-Forwarded-User header
	username := c.Request().Header.Get("X-Forwarded-User")
	if username == "" {
		return c.JSON(403, map[string]interface{}{
			"message": "forbidden",
		})
	}

	users := config.GetUsers()

	// Check if user exists
	if v, ok := users.Users[username]; !ok {
		return c.JSON(403, map[string]interface{}{
			"message": "forbidden",
		})
	} else {
		// return www/user.html with files
		return c.Render(200, "index.html", v)
	}
}

func userFile(c echo.Context) error {
	// Check X-Forwarded-User header
	username := c.Request().Header.Get("X-Forwarded-User")
	if username == "" {
		return c.JSON(403, map[string]interface{}{
			"message": "forbidden",
		})
	}

	users := config.GetUsers()

	// Check if user exists
	if v, ok := users.Users[username]; !ok {
		return c.JSON(403, map[string]interface{}{
			"message": "forbidden",
		})
	} else {
		file := c.Param("file")

		// Check if file is in user's files
		found := false
		for _, f := range v {
			if f == file {
				found = true
				break
			}
		}

		if !found {
			return c.JSON(403, map[string]interface{}{
				"message": "forbidden",
			})
		}

		// Open file
		return c.File(config.GetConfig().FilesBasePath + "/" + file)
	}
}
