package datastructure

import "container/list"

/*
golang中栈的练习与使用
*/

type Stack[T any] struct {
	s *list.List
}

func Constructor[T any]() *Stack[T] {
	return &Stack[T]{
		s: list.New(),
	}
}

func (s *Stack[T]) Push(x T) {
	s.s.PushBack(x)
}

func (s *Stack[T]) Pop() T {
	e := s.s.Back()
	s.s.Remove(e)
	return e.Value.(T)
}

func (s *Stack[T]) Top() T {
	return s.s.Back().Value.(T)
}

func (s *Stack[T]) Empty() bool {
	return s.s.Len() == 0
}
