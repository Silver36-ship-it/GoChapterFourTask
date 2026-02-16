package ChapterFour

import "fmt"

func Task6() {
	var base int
	for {
		fmt.Print("Enter base length (1-10): ")
		fmt.Scan(&base)

		if base >= 1 && base <= 10 {
			break
		}
		fmt.Println("Invalid size. Please stay between 1 and 10.")
	}

	fmt.Println("\nYour Triangle:")

	for row := 1; row <= base; row++ {
		for col := 1; col <= row; col++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
