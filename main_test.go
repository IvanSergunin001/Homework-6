package main

import (
	"testing"
	"github.com/stretchr/testify/assert" 
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) { //кажется готово
	h := generateRandomElements(SIZE)

	assert.Equal(t, SIZE, len(h))
}

func TestMaximum (t *testing.T) {
	g := generateRandomElements(SIZE)
	res := maximum(g)


	table := []struct {
        a   []int
        result int
    }{
		{[]int{1, 2, 3}, 3},
		{[]int{-1, -5, 0}, 0},
		{g, res},
	}

	for _, item := range table {
        h := maximum(item.a)
		assert.Equal(t, item.result, h)
	}

}

