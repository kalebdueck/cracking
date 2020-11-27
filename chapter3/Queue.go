package chapter3

import (
	"fmt"
)

type ele struct {
	name string
	age  int //Really generalizing here...
	next *ele
}

func (e *ele) getAge() int {
	return e.age
}

type queue struct {
	head *ele
}

func initQueue() *queue {
	return &queue{}
}

func (q *queue) Peek() *ele {
	return q.head
}

func (q *queue) isEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}

func (q *queue) Enqueue(name string, age int) {
	ele := &ele{
		name: name,
		age:  age,
	}

	if q.head == nil {
		q.head = ele
	} else {
		current := q.head
		for current.next != nil {
			current = current.next
		}
		current.next = ele
	}
}

func (q *queue) Dequeue() (*ele, error) {

	if q.head == nil {
		return nil, fmt.Errorf("queue is empty")
	}

	q.head = q.head.next

	return q.head, nil
}
