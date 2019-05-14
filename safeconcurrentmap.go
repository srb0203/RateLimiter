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

type mapElement struct {
	firstRequestTime   time.Time
	credits            int
	firstRequestFromIP bool
	timeRemaining      float64
}
