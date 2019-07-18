package main

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
)


// handleExecOnDemandStatus this function execute on demand the check for the given Title
func handleExecOnDemandStatus(c *gin.Context) {
	key := c.Param("key")
	fmt.Printf("\n\n######### start %s", key)
	time.Sleep(60 * time.Second)
	fmt.Printf("\n\n######### end %s", key)
	c.JSON(http.StatusOK, "OK")
}
