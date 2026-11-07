package main

import "sync"

var (
	records = make(map[int]Record)
	nextID  = 1
	mu      sync.Mutex
)
