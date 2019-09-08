package actions

import "github.com/gobuffalo/buffalo"

// HealthCheck default implementation.
func HealthCheck(c buffalo.Context) error {
	return c.Render(200, r.JSON(map[string]string{"message": "Up and running, yo"}))
}
