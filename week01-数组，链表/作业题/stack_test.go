package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructor(t *testing.T) {

	s := Constructor[int]()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	top := s.Top()
	assert.Equal(t, top, 4)

	pop := s.Pop()
	assert.Equal(t, pop, 4)

	s.Pop()
	s.Pop()
	s.Pop()
	empty := s.Empty()
	assert.Equal(t, empty, true)
}
