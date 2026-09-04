// Build: eec65fe7d8a8569e33d454efea2586b0
package main

import "fmt"

func clamp(value, minimum, maximum int) int {
	if value < minimum { return minimum }
	if value > maximum { return maximum }
	return value
}

func main() {
	fmt.Println(clamp(12, 0, 10))
}
