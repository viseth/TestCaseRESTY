package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	// resty.v1
	//"gopkg.in/resty.v1"
	"github.com/go-resty/resty/v2"
)

// RunMonitor func
func RunMonitor(c *gin.Context) {
	postKey := c.PostForm("key")
	monitorEndPoint := "localhost:8050"
	client := resty.New()
	client.SetDebug(true)
	client.SetTimeout(time.Duration(2 * time.Minute))
	//client.SetRetryCount(1)
	resp, err := client.R().
		SetPathParams(map[string]string{
			"key":             postKey,
			"monitorEndPoint": monitorEndPoint,
		}).Get("http://{monitorEndPoint}/v2/{key}")
	if err != nil || resp.StatusCode() != 200 {
		fmt.Printf("\n===> RunMonitor - Failed with error: %s", err)
	} else {
		fmt.Printf("\n===> RunMonitor - Successfully")
	}

	if resp.IsSuccess() {
		statusJSON := fmt.Sprintf("%v", resp)
		fmt.Printf("\nrunMOnitor: %s", statusJSON)
		c.JSON(200, statusJSON)
	} else {
		c.JSON(200, "KO")
	}
}

// Index func
func Index(c *gin.Context) {
	c.Redirect(301, "/layout.html#/keysTiles")
}
