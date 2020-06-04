package main

import (
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	//"github.com/ms-alex/"
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
		path := c.Query("path")
		var body string
		var errMssg error

		if path != "" {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				errMssg = err
			}

			body = string(content)
		}

		c.HTML(http.StatusOK, "pages/path_traversal.gohtml", gin.H{
			"body": body,
			"error": errMssg,
		})
	})

	r.GET("/ssrf", func(c *gin.Context) {
		url := c.Query("url")
		var body string
		var errMssg error

		if url != "" {
			resp, err := http.Get(url)

			if err != nil {
				errMssg = err
			}
			defer resp.Body.Close()
			body_bytes, _ := ioutil.ReadAll(resp.Body)
			body = string(body_bytes)
		}

		c.HTML(http.StatusOK, "pages/ssrf.gohtml", gin.H{
			"body": body,
			"error": errMssg,
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



