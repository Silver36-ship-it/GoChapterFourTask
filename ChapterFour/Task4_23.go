package ChapterFour

import (
	"fmt"
	"math"
)

func Task5() {
	var number int
	largest1 := math.MinInt
	largest2 := math.MinInt

	fmt.Println("Enter 10 integers to find the two largest:")

	for i := 1; i <= 10; i++ {
		fmt.Printf("Enter integer %d: ", i)
		fmt.Scan(&number)

		if number > largest1 {
			largest2 = largest1
			largest1 = number
		} else if number > largest2 {
			largest2 = number
		}
	}

	fmt.Printf("\nWinner (Largest): %d\n", largest1)
	fmt.Printf("Runner-up (Second Largest): %d\n", largest2)
}
