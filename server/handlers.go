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
	if err := c.BindJson(&data); err != nil {
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

	lib.ReadJSON(resp.Body, &country)

	// sql and perform insert, returning the new visit_id
	statement := "insert into visits ( visitor_ip, visitor_country, page, referrer ) values ( $1 $2 $3 $4 ) returning ( visit_id )"

	if res, err := conn.Exec(ctx, statement, country.Ip, country.Country, data.Page, data.Referrer); err != nil {
		fmt.Printf("Error on insert of visit data: %s\n", err)
		return
	}

	id := ""
	if len(res) > 0 { // response
		id = fmt.Sprintf("%d", res[0].visit_id) // 0th row of the response
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"visit_id": fmt.Sprintf("%q", id),
	})
}
