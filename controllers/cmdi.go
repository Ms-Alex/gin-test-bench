package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"os/exec"
)


func CmdInjectionHandler (c *gin.Context) {
	command := c.Query("command")
	var body string
	var errMssg string

	result, err := exec.Command("sh", "-c", command).CombinedOutput()
	if err != nil {
		errMssg = fmt.Sprintf("Exec Err: %v", err)
	}

	body = string(result)

	c.HTML(http.StatusOK, "pages/cmdi.gohtml", gin.H{
		"body": body,
		"error": errMssg,
	})
}