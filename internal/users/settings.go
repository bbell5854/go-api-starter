package users

import (
	"go-api-starter/internal/database"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
)

// UserSettings : Model for GetUserSettings request
type UserSettings struct {
	FirstName string `json:"firstname,omitempty" db:"first_name"`
	LastName  string `json:"lastname,omitempty" db:"last_name"`
	Email     string `json:"email,omitempty" db:"email"`
}

const getUserSettingsQuery = `
SELECT
	um.first_name,
	um.last_name,
	um.email
FROM users u
LEFT JOIN user_meta um ON um.user_id = u.user_id
WHERE 1
	AND u.auth_slug = ?;
`

// GetUserSettings :  fetches user settings from the DB and returns a JSON response
func GetUserSettings(c *fiber.Ctx) {
	// Get slug from request
	userSlug := c.Locals("userSlug")

	var userSettings UserSettings
	err := database.DB.Get(&userSettings, getUserSettingsQuery, userSlug)
	if err != nil {
		log.Println(err)
		c.SendStatus(http.StatusInternalServerError)
		return
	}

	_ = c.JSON(userSettings)
}
