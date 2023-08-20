package main

import (
	gee2 "gee/gee"
	"net/http"
)

func main() {
	r := gee2.New()
	r.GET("/", func(c *gee2.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello/:name", func(c *gee2.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s\n", c.Param("name"), c.Path)
	})
	r.POST("/login", func(c *gee2.Context) {
		c.JSON(http.StatusOK, gee2.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	r.Run(":9999")
}
