package main

// Copyright (C) Jacob Centner, 2023
// BSD 3 Clause Licensed
// See /LICENSE.bsd

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pschlump/dbgo"
)

// command line args

var HostPort = flag.String("hostport", ":9001", "Host/Port to listen on")

//var Dir = flag.String("dir", "./www", "Directory from which to serve static assets")

// database context and connection

var conn *pgxpool.Pool
var ctx context.Context

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
	if len(fns) != 0 {
		fmt.Printf("CLI arguments are not supported: [%s]\n", fns)
		os.Exit(1)
	}

	// router setup
	r := gin.Default()
	// r.Use(static.Serve("/", static.LocalFile(*Dir, false))) // not serving files

	// database connection
	PGXConnect()
	defer PGXDisconnect()

	// ----------------------------------------------------------------------
	// API: /status, /visit,
	// ----------------------------------------------------------------------

	// --- status ---

	r.GET("/api/v1/status", func(c *gin.Context) { // anon inline function defines handler
		c.Status(http.StatusOK)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(http.StatusOK /*200*/, dbgo.SVarI(c))
	})

	// --- visit ---

	r.POST("/api/v1/visit", VisitHandler) // handlers.go

	// ----------------------------------------------------------------------

	r.Run(*HostPort) // listen and serve on HostPort (0.0.0.0:9001)
}
