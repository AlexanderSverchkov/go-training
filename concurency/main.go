package main

import (
	"fmt"
	"math"
	"math/rand"
	randv2 "math/rand/v2"
	"sync"
	"time"
)

func main() {
	channelSize := 10
	chInt := make(chan int)
	chSquare := make(chan int)
	var wg sync.WaitGroup
	go createSliceWithNumbers(chInt, channelSize)
	wg.Add(1)
	go func() {
		defer wg.Done()
		squareNumbers(chInt, chSquare, channelSize)
	}()
	go func() {
		wg.Wait()
		close(chSquare)
	}()
	for res := range chSquare {
		fmt.Println(res)
	}

}

func createSliceWithNumbers(chInt chan int, channelSize int) {
	rand.Seed(time.Now().UnixNano())
	randomInts := make([]int, channelSize)
	for i := 0; i < channelSize; i++ {
		randomInts[i] = randv2.IntN(100)
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
