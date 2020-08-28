package EComApp

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"strings"
)

type BasicAuth struct {
	App         *Application
	ContextName string
}

func (basic *BasicAuth) Use() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			basic.respondWithError(401, "Unauthorized", basic.App, c)
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 || !basic.authenticateUser(pair[0], pair[1], basic.App) {
			basic.respondWithError(401, "Unauthorized", basic.App, c)
			return
		}

		c.Next()
	}
}

func (basic *BasicAuth) respondWithError(code int, message string, app *Application, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}

func (basic *BasicAuth) authenticateUser(username, password string, app *Application) bool {
	if username == "123" && password == "456" {
		return true
	}
	return false
}
