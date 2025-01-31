package queue

import (
	"graph/stack"
)

type Queue struct {
	s1 *stack.Stack
	s2 *stack.Stack
}

func NewQueue() *Queue {
	return &Queue {
		s1: stack.NewStack(),
		s2: stack.NewStack(),
	}
}

func (q *Queue) Push(value int) {
	q.s1.Push(value)
}

func (q *Queue) Pop() (int, bool) {
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			value, _ := q.s1.Pop()
			q.s2.Push(value)
		}
	}
	if !q.s2.IsEmpty() {
		value, _ := q.s2.Pop()
		return value, true
	}
	return 0, false
}

func (q *Queue) IsEmpty() bool {
	return q.s1.IsEmpty() && q.s2.IsEmpty()
}

func (q *Queue) Front() (int, bool) {
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			value, _ := q.s1.Pop()
			q.s2.Push(value)
		}
	}
	if !q.s2.IsEmpty() {
		value := q.s2.Top()
		return value, true
	}
	return 0, false
}

func (q *Queue) Back() (int, bool) {
	if !q.s1.IsEmpty() {
		value := q.s1.Top()
		return value, true
	}
	return 0, false
}
