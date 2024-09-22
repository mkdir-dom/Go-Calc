package main

import "fmt"

func main() {
	var firstnum float64
	var secondnum float64
	var operator string

	resulttimes := times(firstnum, secondnum)

	fmt.Println("Please enter your first number")
	fmt.Scan(&firstnum)

	fmt.Println("Please enter your second number")
	fmt.Scan(&secondnum)

	fmt.Println("Please enter your operator")
	fmt.Scan(&operator)

	if operator == "*" {
		fmt.Println(resulttimes)
	}

}

func times(firstnum, secondnum float64) float64 {
	return firstnum + secondnum
}
