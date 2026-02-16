package ChapterFour

import (
	"fmt"
)

func Maine() {
	var miles, gallons int
	var totalMiles, totalGallons int

	fmt.Println("Mileage Calculator (Enter -1 to quit)")

	for {
		fmt.Println("Enter miles driven:")
		fmt.Scan(&miles)

		if miles == -1 {
			break
		}

		fmt.Println("Enter gallons used:")
		fmt.Scan(&gallons)

		tripMilesPerGallon := float64(miles) / float64(gallons)
		fmt.Printf("Miles per Gallon trip: %.2f\n", tripMilesPerGallon)

		totalMiles += miles
		totalGallons += gallons

		totalMilesPerGallon := float64(totalMiles) / float64(totalGallons)
		fmt.Printf("Combined MilesPerGallon so far: %.2f\n\n", totalMilesPerGallon)
	}
	fmt.Println("Exiting....Drive Safe!")
}
