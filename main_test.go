package main

import (
	"reflect"
	"testing"
)

// TestGenerateRandomElementsZeroSize тестирует генерацию с размером 0
func TestGenerateRandomElementsZeroSize(t *testing.T) {
	result := generateRandomElements(0)

	// Проверяем, что результат — пустой слайс
	if len(result) != 0 {
		t.Errorf("ожидается пустой слайс, но получен слайс с %d элементами", len(result))
	}
}

// TestGenerateRandomElementsNegativeSize тестирует генерацию с отрицательным размером
func TestGenerateRandomElementsNegativeSize(t *testing.T) {
	result := generateRandomElements(-5)

	// Проверяем, что результат — пустой слайс
	if len(result) != 0 {
		t.Errorf("ожидается пустой слайс, но получен слайс с %d элементами", len(result))
	}
}

// TestGenerateRandomElementsPositiveSize тестирует генерацию с положительным размером
func TestGenerateRandomElementsPositiveSize(t *testing.T) {
	size := 10
	result := generateRandomElements(size)

	// Проверяем, что размер слайса совпадает с ожидаемым
	if len(result) != size {
		t.Errorf("ожидается %d элементов, но получен слайс с %d элементами", size, len(result))
	}

	// Проверяем, что все числа положительные (от 1 до 1 000 000)
	for i, value := range result {
		if value < 1 || value > 1_000_000 {
			t.Errorf("элемент [%d] = %d, но ожидается значение от 1 до 1 000 000", i, value)
		}
	}
}

// TestGenerateRandomElementsSingleElement тестирует генерацию с одним элементом
func TestGenerateRandomElementsSingleElement(t *testing.T) {
	result := generateRandomElements(1)

	if len(result) != 1 {
		t.Errorf("ожидается 1 элемент, но получен слайс с %d элементами", len(result))
	}

	// Проверяем, что число положительное
	if result[0] < 1 || result[0] > 1_000_000 {
		t.Errorf("ожидается значение от 1 до 1 000 000, но получено %d", result[0])
	}
}

// TestMaximumEmptySlice тестирует поиск максимума в пустом слайсе
func TestMaximumEmptySlice(t *testing.T) {
	data := []int{}
	result := maximum(data)

	if result != 0 {
		t.Errorf("ожидается 0 для пустого слайса, но получено %d", result)
	}
}

// TestMaximumSingleElement тестирует поиск максимума в слайсе с одним элементом
func TestMaximumSingleElement(t *testing.T) {
	data := []int{42}
	result := maximum(data)

	if result != 42 {
		t.Errorf("ожидается 42, но получено %d", result)
	}
}

// TestMaximumMultipleElements тестирует поиск максимума в слайсе с несколькими элементами
func TestMaximumMultipleElements(t *testing.T) {
	data := []int{3, 7, 2, 9, 5, 1}
	result := maximum(data)

	if result != 9 {
		t.Errorf("ожидается 9, но получено %d", result)
	}
}

// TestMaximumAllSameElements тестирует поиск максимума в слайсе с одинаковыми элементами
func TestMaximumAllSameElements(t *testing.T) {
	data := []int{5, 5, 5, 5, 5}
	result := maximum(data)

	if result != 5 {
		t.Errorf("ожидается 5, но получено %d", result)
	}
}

// TestMaximumSortedAscending тестирует поиск максимума в отсортированном по возрастанию слайсе
func TestMaximumSortedAscending(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := maximum(data)

	if result != 10 {
		t.Errorf("ожидается 10, но получено %d", result)
	}
}

// TestMaximumSortedDescending тестирует поиск максимума в отсортированном по убыванию слайсе
func TestMaximumSortedDescending(t *testing.T) {
	data := []int{100, 90, 80, 70, 60, 50}
	result := maximum(data)

	if result != 100 {
		t.Errorf("ожидается 100, но получено %d", result)
	}
}

// TestMaximumLargeValue тестирует поиск максимума с большим значением
func TestMaximumLargeValue(t *testing.T) {
	data := []int{1, 2, 999999, 4, 5}
	result := maximum(data)

	if result != 999999 {
		t.Errorf("ожидается 999999, но получено %d", result)
	}
}

// TestMaximumVsGenerateRandomElementsConsistency проверяет, что maximum и generateRandomElements работают вместе корректно
func TestMaximumVsGenerateRandomElementsConsistency(t *testing.T) {
	// Генерируем небольшой слайс для быстрого теста
	data := generateRandomElements(100)

	// Находим максимум нашей функцией
	result := maximum(data)

	// Проверяем, что результат не отрицательный и не превышает 1 000 000
	if result < 1 || result > 1_000_000 {
		t.Errorf("максимум %d выходит за допустимые границы (1-1 000 000)", result)
	}

	// Проверяем, что найденный максимум действительно присутствует в слайсе
	found := false
	for _, value := range data {
		if value == result {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("найденное максимальное значение %d отсутствует в слайсе", result)
	}
}

// Helper функция для ручного сравнения результатов single-thread и multi-thread
func compareResults(t *testing.T) {
	data := generateRandomElements(1000)

	singleResult := maximum(data)
	multiResult := maxChunks(data)

	if singleResult != multiResult {
		t.Errorf(
			"результаты не совпадают: single-thread = %d, multi-thread = %d",
			singleResult, multiResult,
		)
	}
}

// TestMaxChunksReturnsCorrectResult проверяет, что maxChunks возвращает правильный результат
func TestMaxChunksReturnsCorrectResult(t *testing.T) {
	// Тест на известных данных
	data := []int{3, 7, 2, 9, 5, 1, 8, 4, 6, 10}
	result := maxChunks(data)

	// Ожидаемый максимум — 10
	if result != 10 {
		t.Errorf("ожидается 10, но получено %d", result)
	}
}

// TestMaxChunksEmptySlice проверяет обработку пустого слайса
func TestMaxChunksEmptySlice(t *testing.T) {
	data := []int{}
	result := maxChunks(data)

	if result != 0 {
		t.Errorf("ожидается 0 для пустого слайса, но получено %d", result)
	}
}

// TestMaxChunksConsistencyWithMaximum проверяет, что maxChunks даёт тот же результат, что и maximum
func TestMaxChunksConsistencyWithMaximum(t *testing.T) {
	data := generateRandomElements(10000)

	singleResult := maximum(data)
	multiResult := maxChunks(data)

	if !reflect.DeepEqual(singleResult, multiResult) {
		t.Errorf(
			"результаты не совпадают: maximum() = %d, maxChunks() = %d",
			singleResult, multiResult,
		)
	}
}
