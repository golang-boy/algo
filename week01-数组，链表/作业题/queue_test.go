package datastructure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeQueue(t *testing.T) {

	q := MakeQueue[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	v := q.Dequeue()
	assert.Equal(t, 1, v)

	v = q.Dequeue()
	assert.Equal(t, 2, v)

	v = q.Dequeue()
	assert.Equal(t, 3, v)

	v = q.Dequeue()
	assert.Equal(t, 4, v)

	e := q.IsEmpty()
	assert.Equal(t, true, e)
}
