package ChapterFour

import "fmt"

func Task9() {
	var num1, num2 int

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	if num1 == num2 {
		fmt.Println(0)
	} else if num1 > num2 {
		fmt.Println(1)
	} else {
		fmt.Println(-1)
	}
}
