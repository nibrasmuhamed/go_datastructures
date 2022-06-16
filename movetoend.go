package main

import "fmt"

func movetoend(array []int, number int) []int {
	i := 0
	j := len(array) - 1

	for i = 0; i < j; i++ {
		if array[j] == number {
			j--
		} else if array[i] == number {
			temp := array[i]
			array[i] = array[j]
			array[j] = temp
		}
	}
	return array
}

func main() {
	arr := []int{12, 21, 12, 2, 12, 3, 78, 90, 9, 6}
	fmt.Println(movetoend(arr, 12))

}
