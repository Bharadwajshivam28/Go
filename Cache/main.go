package main

import (
	"fmt"
)

type Node struct {
val string
Left *Node
Right *Node
}

type Queue struct {
Head *Node
Tail *Node
Length int
}

type Cache struct {
Queue Queue
Hash Hash
}

func NewCache() Cache{
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue{
	head := &Node{}
	tail := &Node{}
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(str string){
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove()
	}else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)
	left := n.Left
	right := n.Right
    left.Right = right
	right.Left = left
	c.Queue.Length

}

type Hash map[string]*Node

func main () {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{"parrot", "avocado", "Dragonfruit", "tree", "potato", "tomato", "tree", "dog"}{
		cache.Check(word)
		cache.Display()
	}
}