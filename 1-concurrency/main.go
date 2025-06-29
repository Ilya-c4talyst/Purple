package main

import (
	"fmt"
	"math/rand"
)

// Создание слайса
func createSlice(lenght int, output chan<- int) {
	defer close(output)
	for range lenght {
		output <- rand.Intn(101)
	}
}

// Возведение в квадрат
func square(input <-chan int, output chan<- int) {
	defer close(output)
	for value := range input {
		output <- value * value
	}
}

func main() {

	// Количество элементов
	lenght := 10
	// Инит необходимых каналов
	createCh := make(chan int, lenght)
	squareCh := make(chan int, lenght)

	// Pipeline
	go createSlice(lenght, createCh)
	go square(createCh, squareCh)

	results := make([]int, lenght)
	i := 0

	for value := range squareCh {

		// Печать результата
		if i == 0 {
			fmt.Print(value)
		} else {
			fmt.Print(" ", value)
		}

		// Сохранение в слайс
		results[i] = value
		i++
	}
}
