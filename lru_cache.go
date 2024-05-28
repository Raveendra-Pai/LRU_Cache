package main

import "fmt"

type LRUCache struct {
	lookup map[string]*Node
	head   *Node
	tail   *Node
	cap    int
	len    int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		lookup: make(map[string]*Node, capacity),
		head:   nil,
		tail:   nil,
		cap:    capacity,
		len:    0,
	}
}

func (c *LRUCache) Get(key string) (any, error) {

	data, ok := c.lookup[key]

	if !ok {
		fmt.Printf("Key %v is not found", key)
		return "", fmt.Errorf("not found")
	}

	if data != c.head {

		data.prev.next = data.next

		if data != c.tail {
			data.next.prev = data.prev
		} else {
			c.tail = c.tail.prev
		}

		c.setHead(data)
	}

	return data.value, nil
}

func (c *LRUCache) updateTail() {
	c.tail = c.tail.prev
	c.tail.next = nil
}

func (c *LRUCache) setHead(node *Node) {
	temp := c.head
	c.head = node
	c.head.next = temp
	temp.prev = c.head
	c.head.prev = nil
}

func (c *LRUCache) Put(k string, v string) {

	if existingNode, ok := c.lookup[k]; ok {
		if existingNode == c.tail {
			c.updateTail()
		}
		c.setHead(existingNode)
		return
	}

	if c.len >= c.cap {
		fmt.Printf("Deleing key: %v from map", c.tail.value)
		delete(c.lookup, c.tail.key)
		c.updateTail()
		c.len--

	}

	newNode := &Node{value: v, key: k, prev: nil, next: nil}
	if c.head == nil {
		c.head = newNode
		c.tail = c.head
	} else {
		c.setHead(newNode)
	}

	c.len++
	c.lookup[k] = newNode
}

func (c *LRUCache) Head() any {
	if c.head != nil {
		return c.head.value
	}
	return ""
}

func (c *LRUCache) Tail() any {
	if c.tail != nil {
		return c.tail.value
	}
	return ""
}
