package datastructure

import "container/list"

/*
golang中队列的练习与使用
*/

type Queue[T any] struct {
	l *list.List
}

func MakeQueue[T any]() *Queue[T] {

	return &Queue[T]{
		l: list.New(),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.l.PushBack(item)
}

func (q *Queue[T]) Dequeue() T {

	if q.l.Len() == 0 {
		return *new(T)
	}
	return q.l.Remove(q.l.Front()).(T)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.l.Len() == 0
}

func (q *Queue[T]) Size() int {
	return q.l.Len()
}

func (q *Queue[T]) Clear() {
	q.l.Init()
}
