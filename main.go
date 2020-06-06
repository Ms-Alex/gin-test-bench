package main

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/Ms-Alex/gin-test-bench/controllers"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	r.Static("/public", "./views/public")

	//Testing
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/home.gohtml", gin.H{
		})
	})

	r.GET("/path-traversal", func(c *gin.Context) {
		controllers.PathTraversalHandler(c)
	})

	r.GET("/ssrf", func(c *gin.Context) {
		controllers.SsrfHandler(c)
	})

	r.GET("/cmdi", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/cmdi.gohtml", gin.H{
		})
	})

	r.GET("/sqli", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/sqli.gohtml", gin.H{
		})
	})

	fmt.Println("Listening on http://localhost:8080")
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}



