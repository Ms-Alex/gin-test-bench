package controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func PathTraversalHandler (c *gin.Context) {
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
}