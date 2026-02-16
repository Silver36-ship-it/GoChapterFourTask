package ChapterFour

import "fmt"

func Task10() {
	var x1, y1, x2, y2 float64

	fmt.Println("Enter coordinates for Point 1 (x1, y1):")
	fmt.Scan(&x1, &y1)

	fmt.Println("Enter coordinates for Point 2 (x2, y2):")
	fmt.Scan(&x2, &y2)

	if x1 == x2 && y1 == y2 {
		fmt.Println("The points are the same location, not a line.")
	} else if x1 == x2 {
		fmt.Println("The points are on a line perpendicular to the X-axis, which is Vertical.")
	} else if y1 == y2 {
		fmt.Println("The points are on a line perpendicular to the Y-axis, which is Horizontal.")
	} else {
		fmt.Println("The points are on a slanted line, which is not perpendicular to any axis.")
	}
}
