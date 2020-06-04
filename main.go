package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	//"github.com/ms-alex/"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	r.Static("/public", "./views/public")

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
		c.HTML(http.StatusOK, "pages/path_traversal.gohtml", gin.H{
		})
	})

	r.GET("/ssrf", func(c *gin.Context) {

		url := c.Query("url")
		var body string
		if url != "" {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			body_bytes, _ := ioutil.ReadAll(resp.Body)
			body = string(body_bytes)
		}

		c.HTML(http.StatusOK, "pages/ssrf.gohtml", gin.H{
			"body": body,
		})
	})

	r.GET("/cmdi", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/cmdi.gohtml", gin.H{
		})
	})

	r.GET("/sqli", func(c *gin.Context) {
		c.HTML(http.StatusOK, "pages/sqli.gohtml", gin.H{
		})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}



