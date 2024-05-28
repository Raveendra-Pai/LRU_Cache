package main

import (
	"fmt"
)

func main() {

	LRUCache := NewLRUCache(3)

	LRUCache.Put("One", "1")
	fmt.Printf("\n 1. Head: %v.... tail: %v\r\n", LRUCache.Head(), LRUCache.Tail())
	LRUCache.Put("two", "2")
	fmt.Printf("\n 2. Head: %v.... tail: %v\r\n", LRUCache.Head(), LRUCache.Tail())
	LRUCache.Put("Three", "3")
	fmt.Printf("\n 3.Head: %v.... tail: %v\r\n", LRUCache.Head(), LRUCache.Tail())
	LRUCache.Put("Four", "4")
	fmt.Printf("\n 4. Head: %v.... tail: %v\r\n", LRUCache.Head(), LRUCache.Tail())

	val, err := LRUCache.Get("two")

	if err == nil {
		fmt.Println("Find:", val)
	}

	fmt.Printf("\n Head: %v.... tail: %v\r\n", LRUCache.Head(), LRUCache.Tail())

	val, err = LRUCache.Get("Four")
	if err == nil {
		fmt.Println("4. Find:", val)
	}

	fmt.Printf("\n Final Head: %v.... tail: %v\r\n", LRUCache.Head(), LRUCache.Tail())

	_, err = LRUCache.Get("One")

	if err != nil {
		fmt.Println(err)
	}

}
