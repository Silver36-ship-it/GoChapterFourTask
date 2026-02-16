package ChapterFour

import "fmt"

type Customer struct {
	AccountNumber    int
	BeginningBalance int
	TotalCharges     int
	TotalCredits     int
	AllowedLimit     int
}

func Task2() {
	for {
		var customer Customer

		fmt.Println("Enter account number (press -1 to quit)")
		fmt.Scan(&customer.AccountNumber)

		if customer.AccountNumber == -1 {
			break
		}

		fmt.Print("Enter Beginning Balance:")
		fmt.Scan(&customer.BeginningBalance)

		fmt.Print("Enter total charges for the month:")
		fmt.Scan(&customer.TotalCharges)

		fmt.Print("Enter total Credits applied:")
		fmt.Scan(&customer.TotalCredits)

		fmt.Print("Enter Credit limit: ")
		fmt.Scan(&customer.AllowedLimit)

		newBalance := customer.BeginningBalance + customer.TotalCharges - customer.TotalCredits

		fmt.Printf("New Balance for account number %d is : %d\n", customer.AccountNumber, newBalance)

		if newBalance > customer.AllowedLimit {
			fmt.Println("Credit limit exceeded!")
		}
	}
	fmt.Println("Exiting...")
}
