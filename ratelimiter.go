package main

import (
	"fmt"
	"net/http"
	"time"
)

var globalMap = safeConcurrentMap{value: make(map[string]mapElement)}

func limitExceeded(ipAddr string, numberOfRequests int, timeLimit float64) bool {
	fmt.Println("ipAddr: " + ipAddr)

	firstRequestTime := globalMap.get(ipAddr).firstRequestTime
	duration := time.Now().Sub(firstRequestTime)

	if duration.Seconds() <= timeLimit {
		var m mapElement
		m.firstRequestFromIP = true
		m.credits = globalMap.get(ipAddr).credits - 1
		m.firstRequestTime = globalMap.get(ipAddr).firstRequestTime
		m.timeRemaining = timeLimit - duration.Seconds()
		globalMap.set(ipAddr, m)
	} else {
		var m mapElement
		m.credits = numberOfRequests - 1 //subtract current request served
		m.firstRequestTime = time.Now()
		m.firstRequestFromIP = true
		m.timeRemaining = timeLimit - duration.Seconds()
		globalMap.set(ipAddr, m)
	}

	if globalMap.get(ipAddr).credits < 0 {
		return true
	}
	return false
}

//GetUserIPAddress : find user IP address from the request
func getUserIPAddress(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func initializeIPInMap(ipAddr string, t time.Time, numberOfRequests int) {
	var me mapElement
	me.firstRequestTime = t
	me.credits = numberOfRequests
	globalMap.set(ipAddr, me)
}

func rateLimit(h http.HandlerFunc, numberOfRequests int, timeLimit float64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ipAddr := getUserIPAddress(r)
		t := time.Now()

		//First request from this IP, initialize variables
		if !globalMap.get(ipAddr).firstRequestFromIP {
			initializeIPInMap(ipAddr, t, numberOfRequests)
		}

		if limitExceeded(ipAddr, numberOfRequests, timeLimit) {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string("Too many requests, please try again in %f seconds"),
				globalMap.get(ipAddr).timeRemaining)
		} else {
			if h == nil {
				http.DefaultServeMux.ServeHTTP(w, r)
				return
			}
			h.ServeHTTP(w, r)
		}
	}
}
