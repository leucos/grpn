package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushPop(t *testing.T) {
	var s Stack

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.Equal(t, 3.0, s.Pop())
	assert.Equal(t, 2.0, s.Pop())
	assert.Equal(t, 1.0, s.Pop())
}

func TestPopNil(t *testing.T) {
	var s Stack

	assert.Nil(t, s.Pop())
}

func TestDup(t *testing.T) {
	var s Stack

	s.Push(1)
	s.Dup()

	assert.Equal(t, 1.0, s.Pop())
	assert.Equal(t, 1.0, s.Pop())
}

func TestSwap(t *testing.T) {
	var s Stack

	s.Push(1).Push(2).Swap()

	assert.Equal(t, 1.0, s.Pop())
	assert.Equal(t, 2.0, s.Pop())
}
func TestEmpty(t *testing.T) {
	var s Stack

	assert.True(t, s.IsEmpty())

	s.Push(1)
	assert.False(t, s.IsEmpty())

}
