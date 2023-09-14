package main

import (
	"fmt"
	"log"
)

type Node struct {
	info interface{}
	next *Node
}

type ListNode struct {
	head *Node
}

func (l *ListNode) Insert(d interface{}) {
	list := &Node{info: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}

func Show(l *ListNode) {
	p := l.head
	for p != nil {
		fmt.Printf("--> %v", p.info)
		p = p.next
	}
}

func main() {
	sl := ListNode{}
	var input int
	fmt.Println("Enter a value:")
	for i := 0; i < 10; i++ {
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err.Error())
		}
		sl.Insert(input)
		Show(&sl)
	}
}
