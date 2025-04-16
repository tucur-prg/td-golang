package main

import (
	"fmt"
	"sync"
)

func main() {
	data := Cache{}

	go data.Set("1", "aaa")
	go data.Set("2", "bbb")
	go data.Set("3", "ccc")
	go data.Set("4", "ddd")

	fmt.Println(data.Get("3"))
}

type Cache struct {
	mu sync.RWMutex
	store map[string]string
}

func (s *Cache) Get(key string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.store[key]
}

func (s *Cache) Set(key, val string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.store[key] = val
}
