
package main

import (
	"fmt"
	"strings"
	"net/http"
	"os"
	"net"
	"github.com/gin-gonic/gin"
	"time"
	"gopkg.in/urfave/cli.v1"
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
	app.Name = "webserver"
	app.Usage = "fancy server for connecting to a webserver keystore"
	app.Version = version
	app.Compiled = time.Now()
	app.Action = func(c *cli.Context) error {
		port = c.GlobalString("port")
		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		r.GET("/", Index)
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.POST("/runMonitor", RunMonitor)
		r.Static("/web", "./web")
		fmt.Printf("webserver-server (v.%s) running on http://%s:%s\n", version, GetLocalIP(), port)
		s := &http.Server{
			Addr:           ":"+port,
			Handler:        r,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()				
		return nil
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "9090",
			Usage: "port to use to listen",
		},
	}
	app.Run(os.Args)	
}
