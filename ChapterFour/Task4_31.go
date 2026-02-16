package ChapterFour

import "fmt"

func Task7() {
	var number int

	for {
		fmt.Print("Enter a five-digit integer: ")
		fmt.Scan(&number)

		if number >= 10000 && number <= 99999 {
			break
		}
		fmt.Println("Error: Number must be exactly five digits long.")
	}

	digit1 := number / 10000
	digit2 := (number / 1000) % 10
	digit4 := (number / 10) % 10
	digit5 := number % 10

	if digit1 == digit5 && digit2 == digit4 {
		fmt.Printf("%d is a palindrome!\n", number)
	} else {
		fmt.Printf("%d is NOT a palindrome.\n", number)
	}
}
