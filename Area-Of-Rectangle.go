// finding the area of a rectangle
package main

import (
	"fmt"
)

func main() {
    var length, width float64

    // Prompt user for input
    fmt.Println("Enter the length of the rectangle:")
    fmt.Scanln(&length)
    fmt.Println("Enter the width of the rectangle:")
    fmt.Scanln(&width)

    // Calculate area
    area := length * width

    // Output result
    fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
