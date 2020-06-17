package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnvalidatedRedirectHandler (c *gin.Context) {
	redirectUrl := c.Query("url")
	fmt.Println("Redirect is ", redirectUrl)

	if redirectUrl != "" {
		c.Redirect(http.StatusMovedPermanently, redirectUrl)
	}

	c.HTML(http.StatusOK, "pages/unvalidated_redirect.gohtml", gin.H{})
}