package context

import (
	"net/http"
	"testing"
)

func TestContext(t *testing.T) {
	gee := New()

	gee.GET("/favicon.ico", func(c *Context) {
		c.Data(http.StatusOK, nil)
	})
	gee.GET("/", func(c *Context) {
		c.JSON(http.StatusOK, map[string]any{
			"hello": "world",
		})
	})
	gee.GET("/test", func(c *Context) {
		c.String(http.StatusOK, "test")
	})

	gee.GET("/query", func(c *Context) {
		c.JSON(http.StatusOK, map[string]any{
			"name": c.Query("name"),
		})
	})

	gee.Run()
}
