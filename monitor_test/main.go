package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"net"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

var version string
var port string

// GetLocalIP returns the local ip address
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "localhost"
	}
	bestIP := "localhost"
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil && (strings.Contains(ipnet.IP.String(), "192.168.1") || strings.Contains(ipnet.IP.String(), "192.168")) {
				return ipnet.IP.String()
			}
		}
	}
	return bestIP
}

func main() {
	app := cli.NewApp()
	app.Name = "monitor"
	app.Usage = "fancy monitor"
	app.Version = version
	app.Compiled = time.Now()
	app.Action = func(c *cli.Context) error {
		port = c.GlobalString("port")
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		r.GET("/v2/:key", handleExecOnDemandStatus)
		fmt.Printf("AMi-webserver (v.%s) running on http://%s:%s\n", version, GetLocalIP(), port)
		r.Run(":" + port) // listen and serve on 0.0.0.0:8080
		return nil
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "8050",
			Usage: "port to use to listen",
		},
	}
	app.Run(os.Args)

}
