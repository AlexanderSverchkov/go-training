package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sync"
)

func main() {
	channelSize := 10
	chInt := make(chan int)
	chSquare := make(chan int)
	var wg sync.WaitGroup
	go createSliceWithNumbers(chInt, channelSize)
	wg.Add(1)
	go func() {
		squareNumbers(chInt, chSquare, channelSize)
		wg.Done()
	}()
	for i := 0; i < channelSize; i++ {
		fmt.Println(<-chSquare)
	}
	wg.Wait()
	close(chSquare)
}

func createSliceWithNumbers(chInt chan int, channelSize int) {
	randomInts := make([]int, channelSize)
	for i := 0; i < channelSize; i++ {
		randomInts[i] = rand.IntN(100)
	}
	for i := 0; i < channelSize; i++ {
		chInt <- randomInts[i]
	}
}

func squareNumbers(chInt chan int, chSquare chan int, channelSize int) {
	for i := 0; i < channelSize; i++ {
		val := <-chInt
		chSquare <- int(math.Pow(float64(val), 2))
	}
	close(chInt)
}
