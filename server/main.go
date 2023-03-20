package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/static"
	"github.com/pschlump/dbgo"
)

var HostPort = flag.String("hostport", ":9001", "Host/Port to listen on")

func main() {

	// ----------------------------------------------------------------------
	// CLI Arguments
	// ----------------------------------------------------------------------
	
	// help message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "stats-tracker: Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	fns := flag.Args()
	if lens(fns) != 0 {
		fmt.Printf("CLI arguments are not supported: [%s]\n", fns)
		os.Exit(1)
	}

	// router setup
	r := gin.Default()
	// r.Use(static.Serve("/", static.LocalFile(*Dir, false)	// not serving files

	
	// ----------------------------------------------------------------------
	// API: /status, 
	// ----------------------------------------------------------------------
	
	r.GET("/api/v1/status", func(c *gin.Context) { // anon inline function
		c.Status(http.StatusOk)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(http.StatusOk, dbgo.SVarI(c))
	})
	
	// ----------------------------------------------------------------------
	r.Run(*HostPort) // listen and serve on HostPort (0.0.0.0:9001)
}
