package main

import (
	"fmt"
)

const SIZE = 5

type Node struct {
	val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int // basically a linklist structure where we are representing in
	//queue type where there is head and there is tail
	//new elemnt coonected to head deltion one connected to tail

}

type Cache struct {
	Queue Queue

	Hash Hash
}

type Hash map[string]*Node

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}

}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

func (c *Cache) check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{val: str}
	}
	c.Add(node)
	c.Hash[str] = node

}

func (c *Cache) Remove(n *Node) *Node {

	fmt.Printf("remove : %s\n", n.val)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length -= 1
	delete(c.Hash, n.val)
	return n

}

func (c *Cache) Add(n *Node) {
	fmt.Printf("Adding : %s\n", n.val)
	temp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = temp
	temp.Left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()

}

func (q *Queue) Display() {

	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.val)
		if i < q.Length-1 {
			fmt.Print("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("START CACHE")
	cache := NewCache()
	for _, word := range []string{"book1", "book2", "book3"} {
		cache.check(word)
		cache.Display()
	}
}
