package main

import (
	"testing"
	"github.com/stretchr/testify/assert" 
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	h := generateRandomElements(SIZE)

	assert.Equal(t, SIZE, len(h))
}

func TestMaximum (t *testing.T) {


	table := []struct {
        a   []int
        result int
    }{
		{[]int{1, 2, 3}, 3},
		{[]int{-1, -5, 0}, 0},
		{[]int{7, 7, 7}, 7},
		{[]int{1, 5, 3, 9, 2}, 9},
	}

	for _, item := range table {
        h := maximum(item.a)
		assert.Equal(t, item.result, h)
	}

}

