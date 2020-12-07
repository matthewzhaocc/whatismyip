package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
)

// Gets the IP I guess
func WhatIsMyIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Header.Get("X-FORWARDED-FOR"))
	LoggingContent := "Just received request from " + r.Header.Get("X-FORWARDED-FOR")
	log.Println(LoggingContent)
}

// Starts the Webserver
func main() {
	http.HandleFunc("/", WhatIsMyIP)
	log.Println("Started Listening on :6443")
	http.ListenAndServe(":6443", nil)
}
