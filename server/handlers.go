package main

// Copyright (C) Jacob Centner, 2023
// BSD 3 Clause Licensed
// See /LICENSE.bsd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// VisitHandler inserts a new row into the visits table
// r.POST("/api/v1/visit", VisitHandler)

func VisitHandler(c *gin.Context) {

	type VisitData struct {
		Page     string `json:"page"`
		Referrer string `json:"referrer"`
	}

	type CountryData struct {
		Ip      string `json:"ip"`
		Country string `json:"country"`
	}

	// get json data from api call
	var data VisitData
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest /*400*/, err)
	}

	// get client IP
	ip := c.Request.Header.Get("X-Forwarded-For")

	// get country for IP
	var country CountryData
	var request_string = "https://api.country.is/" + ip

	resp, err := http.Get(request_string)
	if err != nil {
		fmt.Printf("Country API encountered an error: %s, using XX as backup.\n", err)
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	ReadJSON(resp.Body, &country)

	// sql and perform insert, returning the new visit_id
	statement := "insert into visits ( visitor_ip, visitor_country, page, referrer ) values ( $1 $2 $3 $4 ) returning ( visit_id )"

	res, err := conn.Exec(ctx, statement, country.Ip, country.Country, data.Page, data.Referrer)
	if err != nil {
		fmt.Printf("Error on insert of visit data: %s\n", err)
		return
	}

	// check response for visit_id
	defer res.Close() // ensure not using up connections
	var id string
	res.Next() // get next row in response
	if err = res.Scan(&id); err != nil {
		fmt.Printf("Error in reading response from insert query (VisitHandler): %s\n", err)
	}

	// check conn lost
	if err = res.Err(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"visit_id": fmt.Sprintf("%q", id),
	})
}
