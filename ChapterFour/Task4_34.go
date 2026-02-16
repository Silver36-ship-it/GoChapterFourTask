package ChapterFour

import "fmt"

func Task8() {
	var target, input, sum int

	fmt.Print("Enter the target sum value: ")
	fmt.Scan(&target)

	fmt.Printf("Good one! Now enter integers to reach %d.\n", target)

	for sum < target {
		fmt.Printf("Current sum is %d. Enter next number: ", sum)
		fmt.Scan(&input)

		sum += input
	}

	fmt.Printf("\nTarget reached or passed!\nFinal Sum is %d\n", sum)
}
