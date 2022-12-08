package utils

import (
	"errors"
	"sync"
)

type stack struct {
	lock sync.Mutex
	s    []int
}

func NewStack() *stack {
	return &stack{sync.Mutex{}, make([]int, 0)}
}

func (s *stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)

	if l == 0 {
		return 0, errors.New("Empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]

	return res, nil
}

type StringStack struct {
	lock sync.Mutex
	s    []string
}

func NewStringStack() *StringStack {
	return &StringStack{sync.Mutex{}, make([]string, 0)}
}

func (s *StringStack) Push(v string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *StringStack) Pop() (string, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)

	if l == 0 {
		return "", errors.New("Empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]

	return res, nil
}
