package ChapterFour

import "fmt"

func Task3() {
	const ceiling = 30000.0
	const lowRate = 0.15
	const highRate = 0.20

	for i := 1; i <= 3; i++ {
		var name string
		var earnings float64

		fmt.Printf("\nCitizen %d:\n", i)
		fmt.Print("Enter name: ")
		fmt.Scan(&name)
		fmt.Print("Enter annual earnings: ")
		fmt.Scan(&earnings)

		var totalTax float64

		if earnings <= ceiling {
			totalTax = earnings * lowRate
		} else {
			taxOnBase := ceiling * lowRate
			taxOnExcess := (earnings - ceiling) * highRate
			totalTax = taxOnBase + taxOnExcess
		}

		fmt.Printf("Total tax for %s: $%.2f\n", name, totalTax)
	}
}
