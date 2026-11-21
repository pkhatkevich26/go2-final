package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	randomData := make([]int, size)
	for i := 0; i < size; i++ {
		randomData[i] = rand.Intn(1_000_000) + 1
	}

	return randomData
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkMaxima := make([]int, CHUNKS)
	var waitGroup sync.WaitGroup
	waitGroup.Add(CHUNKS)

	chunkSize := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		startIndex := i * chunkSize

		var endIndex int
		if i == CHUNKS-1 {
			endIndex = len(data)
		} else {
			endIndex = startIndex + chunkSize
		}

		go func(index int, slice []int) {
			chunkMax := maximum(slice)
			chunkMaxima[index] = chunkMax
			waitGroup.Done()
		}(i, data[startIndex:endIndex])
	}

	waitGroup.Wait()
	return maximum(chunkMaxima)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	rand.Seed(time.Now().UnixNano())

	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	startTime := time.Now()
	maxSingle := maximum(data)
	elapsedSingle := time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxSingle, elapsedSingle)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	startTime = time.Now()
	maxMulti := maxChunks(data)
	elapsedMulti := time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxMulti, elapsedMulti)
}
