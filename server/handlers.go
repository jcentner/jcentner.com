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
		Page     string `json:"page" binding:"required"`
		Referrer string `json:"referrer" binding:"required"`
	}

	type CountryData struct {
		Ip      string `json:"ip"`
		Country string `json:"country"`
	}

	// get json data from api call
	var data VisitData
	if err := c.BindJSON(&data); err != nil {
		c.AbortWithError(http.StatusNotAcceptable /*406*/, err)
		return
	}

	// get client IP
	ip := c.Request.Header.Get("X-Forwarded-For")

	// get country for IP
	var country CountryData
	var request_string = "https://api.country.is/" + ip

	resp, err := http.Get(request_string)
	if err != nil {
		fmt.Printf("Country API encountered an error: %s, using XX as backup.\n", err)
		country.Country = "XX"
		/* TODO test country backup*/
	}

	defer resp.Body.Close()

	if err = ReadJSON(resp.Body, &country); err != nil {
		fmt.Printf("ReadJSON error: %s, using IP/Country backup\n", err)
		country.Ip = ip
		country.Country = "XX"
	}

	// sql and perform insert, returning the new visit_id
	statement := "insert into visits ( visitor_ip, visitor_country, page, referrer ) values ( $1, $2, $3, $4 ) returning ( visit_id )"

	row := conn.QueryRow(ctx, statement, country.Ip, country.Country, data.Page, data.Referrer)

	// check response for visit_id
	// PGX pool closes connection on row.Scan
	var id string
	if err = row.Scan(&id); err != nil {
		fmt.Printf("Error in reading response from insert query (VisitHandler): %s\n", err)
		c.AbortWithError(http.StatusBadRequest /*400*/, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   "success",
		"visit_id": fmt.Sprintf("%q", id),
	})
}

// SocialclickHandler inserts a new row into the Socialclick table
// r.POST("/api/v1/socialclick", SocialclickHandler

func SocialclickHandler(c *gin.Context) {

	// get client IP
	ip := c.Request.Header.Get("X-Forwarded-For")

	//sql and perform insert, returning the new socialclick_id
	statement := "insert into socialclicks ( visitor_ip ) values ( $1 ) returning ( socialclick_id )"

	row := conn.QueryRow(ctx, statement, ip)

	// check response for id
	// PGX pool closes connection on row.Scan
	var id string
	if err = row.Scan(&id); err != nil {
		fmt.Printf("Error in reading response from insert query (SocialclickHandler): %s\n", err)
		c.AbortWithError(http.StatusBadRequest /*400*/, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":         "success",
		"socialclick_id": fmt.Sprintf("%q", id),
	})
}
