package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SsrfHandler (c *gin.Context) {
	url := c.Query("url")
	var body string
	var errMssg string

	defer func() {
		if err := recover(); err != nil {
			errMssg = fmt.Sprintf("Panic Err: %v", err)
			renderPage(c, body, errMssg)
		}
	}()

	if url != "" {
		resp, err := http.Get(url)
		defer resp.Body.Close()

		if err != nil {
			errMssg = fmt.Sprintf("Get Err: %v", err)
		}

		body_bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errMssg = fmt.Sprintf("%s \nRead Err: %v", errMssg, err)
		}

		body = string(body_bytes)

	}

	renderPage(c, body, errMssg)
}

func renderPage (c *gin.Context, body string, errMssg string) {
	c.HTML(http.StatusOK, "pages/ssrf.gohtml", gin.H{
		"body": body,
		"error": errMssg,
	})
}