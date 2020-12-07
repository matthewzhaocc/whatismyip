package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
)

// Gets the IP I guess
func WhatIsMyIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.RemoteAddr)
	LoggingContent := "Just received request from " + r.RemoteAddr
	log.Println(LoggingContent)
}

// Starts the Webserver
func main() {
	http.HandleFunc("/", WhatIsMyIP)
	log.Println("Started Listening on :6443")
	http.ListenAndServe(":6443", nil)
}
