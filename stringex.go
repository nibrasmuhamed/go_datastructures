package main

import "fmt"

func stringsort(word string) []int {
	var array [26]int
	var j int
	for _, eachWord := range word {
		if eachWord >= 'a' && eachWord <= 'z' {
			j = int(eachWord) - 'a'
			array[j]++
		}
	}
	return array[:]
}

func printVal(array []int) {
	for i := 0; i < len(array); i++ {
		fmt.Println(string(i+'a'), "is ", array[i])
	}
}

func main() {
	anns := stringsort("hello")
	printVal(anns)
}
