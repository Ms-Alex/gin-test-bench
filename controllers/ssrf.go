package controllers

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func SsrfHandler (c *gin.Context) {
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
}