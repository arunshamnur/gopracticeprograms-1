package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	data int
}

type DL struct {
	head *Node
}

func (d *DL) insert(data int) {
	node := Node{
		data: data,
		prev: nil,
		next: nil,
	}
	if d.head != nil {
		current := d.head
		for current.next != nil {
			current = current.next
		}
		current.next = &node
		node.prev = current
	} else {
		d.head = &node
	}
}

func (d *DL) reverse() {
	var next *Node
	var prev *Node
	current := d.head
	for current != nil {
		next, current.next = current.next, prev
		prev, current = current, next
	}
	d.head = prev
}

func (d *DL) print() {
	if d.head != nil {
		current := d.head
		for current != nil {
			fmt.Printf("%v -> ", current.data)
			current = current.next
		}
		fmt.Println()
	}
}

func main() {
	dl := DL{}
	dl.insert(1)
	dl.insert(2)
	dl.insert(3)
	dl.insert(4)
	dl.insert(5)
	dl.print()
	dl.reverse()
	dl.print()
}
