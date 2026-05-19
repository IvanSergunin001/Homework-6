package main

import (
	"fmt"
	"math/rand"
	"sync"
	"slices"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

var mu sync.Mutex
var wg sync.WaitGroup

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int { //кажется готово
	// ваш код здесь
	if size == 0 {
		return []int{}
	}

	numbers := make([]int, size)

	for i := 0; i < size; i++ {
		numbers[i] = rand.Int() 
	}

	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int { //кажется готово
	// ваш код здесь
	if len(data) == 0 {
		return 0
	} else if len(data) == 1 {
		return 1
	}

	maxNumber := slices.Max(data)

	return maxNumber
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int { //по моей логике это должно работать
	// ваш код здесь
	slicePieceLen := len(data) / CHUNKS
	maxChunksSlice := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++{
		wg.Add(1)
		go func (i int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			startIndex := i * slicePieceLen
			endIndex := startIndex + slicePieceLen
			s := data[startIndex:endIndex]
			maxNumberInChunk := slices.Max(s) //
			maxChunksSlice = append(maxChunksSlice, maxNumberInChunk)
		} (i)
	}
	wg.Wait()
	
	maxNumber := slices.Max(maxChunksSlice)

	return maxNumber
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(numbers)
	end := time.Now()
	timeTotal := end.Sub(start)
    elapsed := timeTotal.Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	// ваш код здесь
	start = time.Now()
	max = maxChunks(numbers)
	end = time.Now()
	timeTotal = end.Sub(start)
	elapsed = timeTotal.Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
