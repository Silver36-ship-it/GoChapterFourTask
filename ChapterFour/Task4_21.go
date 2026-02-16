package ChapterFour

import (
	"fmt"
	"math"
)

func Task4() {
	var number int
	largest := math.MinInt

	fmt.Println("Enter 10 integers:")

	for counter := 1; counter <= 10; counter++ {
		fmt.Printf("Enter integer %d: ", counter)
		fmt.Scan(&number)

		if number > largest {
			largest = number
		}
	}

	fmt.Printf("\nThe largest integer is: %d\n", largest)
}
