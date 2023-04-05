package main

// Copyright (C) Jacob Centner, 2023
// BSD 3 Clause Licensed
// See /LICENSE.bsd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

// ReadJSON: read JSON and unmarshal into a go struct

func ReadJSON(body io.ReadCloser, gostruct interface{}) (err error) {
	buffer, err := ioutil.Readall(body)
	if err != nil {
		fmt.Printf("ReadJSON encountered error: %s\n", err)
		return
	}
	err = json.Unmarshal(buffer, gostruct)
	return
}

// PGXConnect populates "conn" global that can connect to a postgres instance
// Requires DATABASE_URL env var set
// Requires "conn" and "ctx" globals
// var conn *pgxpool.Pool
// var ctx context.Context

func PGXConnect() {
	ctx = context.Background()
	conurl := os.Getenv("DATABASE_URL")

	var err error
	if conn, err = pgxpool.Connect(ctx, conurl); err != nil {
		fmt.Printf("Unable to connect to %s; error: %v\n", conurl, err)
		os.Exit(1)
	}
}

//PGXDisconnect closes "conn" global connection

func PGXDisconnect() {
	conn.Close()
}
