// Finding two numbers in an array which has sum of 10

package main

import "fmt"

func findelement(array [10]int, result int) (int, int) {
	var i = 0
	var j = i + 1
	for i = 0; i < len(array); i++ {
		for j = i + 1; j < len(array); j++ {
			if array[i]+array[j] == 10 {
				return array[i], array[j]
			}
		}
	}
	return i, j
}

func main() {
	var arr [10]int = [10]int{12, 21, 7, 2, 32, 3, 78, 90, 9, 6}
	var a, b = findelement(arr, 10)
	fmt.Printf("first element is %d and second element is %d", a, b)

}
