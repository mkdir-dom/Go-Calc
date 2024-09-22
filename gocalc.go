package main

// this still doesn't work

import "fmt"

func main() {
	var firstnum float64
	var secondnum float64
	var operator string

	resulttimes := times(firstnum, secondnum)
	resultdivide := divide(firstnum, secondnum)
	resultadd := add(firstnum, secondnum)
	resultsubtract := subtract(firstnum, secondnum)

	fmt.Println("Please enter your first number")
	fmt.Scan(&firstnum)

	fmt.Println("Please enter your second number")
	fmt.Scan(&secondnum)

	fmt.Println("Please enter your operator")
	fmt.Scan(&operator)

	if operator == "*" {
		fmt.Println(resulttimes)
	} else if operator == "/" {
		fmt.Println(resultdivide)
	} else if operator == "+" {
		fmt.Println(resultadd)
	} else if operator == "-" {
		fmt.Println(resultsubtract)
	} else {
		fmt.Println("Invalid entry")
	}

}

func times(firstnum, secondnum float64) float64 {
	return (firstnum * secondnum)
}

func divide(firstnum, secondnum float64) float64 {
	return (firstnum / secondnum)
}

func add(firstnum, secondnum float64) float64 {
	return (firstnum + secondnum)
}

func subtract(firstnum, secondnum float64) float64 {
	return (firstnum - secondnum)
}
