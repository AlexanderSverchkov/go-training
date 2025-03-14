package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

func main() {
	channelSize := 10
	chInt := make(chan int)
	chSquare := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		createSliceWithNumbers(chInt, channelSize)
		wg.Done()
		close(chInt)
	}()

	go func() {
		defer wg.Done()
		squareNumbers(chInt, chSquare, channelSize)
	}()
	go func() {
		wg.Wait()
		close(chSquare)
	}()

	go func() {
		defer wg.Done()
		printChSquare(chSquare)
	}()

	//wg.Wait()
	//close(chSquare)

}

func createSliceWithNumbers(chInt chan int, channelSize int) {
	SEED := time.Now().UnixNano()       // Dynamic seed based on current time
	r := rand.New(rand.NewSource(SEED)) // Create rand.Rand instance outside the loop
	randomInts := make([]int, channelSize)
	for i := 0; i < channelSize; i++ {
		randomInts[i] = r.Intn(100)
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
	//close(chInt)
}
func printChSquare(chSquare chan int) {
	for res := range chSquare {
		fmt.Println(res)
	}
}
