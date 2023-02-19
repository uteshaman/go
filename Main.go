package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(value interface{}) {
	node := &Node{value, s.top}
	s.top = node
	s.size++
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	value := s.top.value
	s.top = s.top.next
	s.size--
	return value, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return s.top.value, nil
}

func (s *Stack) Clear() {
	s.top = nil
	s.size = 0
}

func (s *Stack) Contains(value interface{}) bool {
	current := s.top
	for current != nil {
		if current.value == value {
			return true
		}
		current = current.next
	}
	return false
}

func (s *Stack) Increment() {
	current := s.top
	for current != nil {
		current.value = current.value.(int) + 1
		current = current.next
	}
}

func (s *Stack) Print() {
	current := s.top
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

func (s *Stack) PrintReverse() {
	values := make([]interface{}, 0, s.size)
	current := s.top
	for current != nil {
		values = append(values, current.value)
		current = current.next
	}
	for i := len(values) - 1; i >= 0; i-- {
		fmt.Println(values[i])
	}
}
