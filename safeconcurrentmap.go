package main

import (
	"sync"
	"time"
)

//Safe map so that only one client can access it a time
type safeConcurrentMap struct {
	value map[string]mapElement
	mux   sync.Mutex
}

//set the map element by first locking the map
func (m *safeConcurrentMap) set(key string, val mapElement) {
	m.mux.Lock()
	m.value[key] = val
	m.mux.Unlock()
}

func (m *safeConcurrentMap) get(key string) mapElement {
	m.mux.Lock()
	defer m.mux.Unlock()
	return m.value[key]
}

//Map element to store values related to rate limitter
type mapElement struct {
	firstRequestTime time.Time
	credits          int
	ipAlreadySeen    bool
	timeRemaining    float64
}
