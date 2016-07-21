package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Print 1 to 10, 10 times with a random delay from 0 - 250ms
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	defer timeChange(time.Now())
	for i := 0; i < 10; i++ {
		go f(i) //takes 3.4 seconds
		//	f(i) //takes 12 seconds
	}

	var input string
	fmt.Scanln(&input)
}

func timeChange(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}
