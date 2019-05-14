package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func makeRequest(t *testing.T, ts *httptest.Server) string {

	res, err := http.Get(ts.URL)
	fmt.Println(ts.URL)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func TestMyHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rateLimit(Index)))
	defer ts.Close()

	for i := 0; i < 2; i++ {

		body := makeRequest(t, ts)
		expected := "Hello World example for GET request.\n"

		if string(body) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				string(body), expected)
		}
	}

	for i := 0; i < 2; i++ {

		body := makeRequest(t, ts)
		expected := "Too many requests, please try again in"

		if !strings.Contains(string(body), expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				string(body), expected)
		}
	}

	time.Sleep(10 * time.Second)

	for i := 0; i < 2; i++ {

		body := makeRequest(t, ts)
		expected := "Hello World example for GET request.\n"

		if string(body) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				string(body), expected)
		}
	}
}