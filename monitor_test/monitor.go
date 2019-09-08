package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// handleExecOnDemandStatus this function execute on demand the check for the given Title
func handleExecOnDemandStatus(c *gin.Context) {
	key := c.Param("key")
	fmt.Printf("\n\n######### start %s", key)
	time.Sleep(90 * time.Second)
	fmt.Printf("\n\n######### end %s", key)
	c.JSON(http.StatusOK, "OK")
}
