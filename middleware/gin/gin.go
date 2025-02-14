package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	s := gin.Default()

	s.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "who are you?")
	})

	s.GET("/panic", func(c *gin.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	s.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//动态参数
	//http://localhost:8080/user/tutu
	s.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	//参数解析
	//http://localhost:8080/user2/tutu?action=do
	s.GET("/user2/:name", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Query("action")
		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
		})
	})

	//POST
	s.POST("/user3", func(c *gin.Context) {
		name := c.PostForm("name")
		action := c.PostForm("action")

		c.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
		})
	})

	//POST&GET
	s.POST("/user4", func(c *gin.Context) {
		id := c.Query("id")
		name := c.PostForm("name")
		action := c.PostForm("action")
		c.JSON(http.StatusOK, gin.H{
			"id":     id,
			"name":   name,
			"action": action,
		})
	})

	//REDIRECT
	s.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	//ROUTE GROUP
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	v1 := s.Group("v1")
	v1.GET("/user", defaultHandler)
	v2 := s.Group("v2")
	v2.GET("/user", defaultHandler)

	//FILE UPLOAD
	s.POST("/file", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
		}
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

	//NODE
	s.Use(gin.Logger())
	s.Use(gin.Recovery())
	s.Use(func(c *gin.Context) {
		cur := time.Now()
		c.Set("hello", "world")
		c.Next()
		elapsed := time.Since(cur)
		fmt.Println("middleware function call, elapsed: ", elapsed)
	})

	s.Run(":8080")
}
