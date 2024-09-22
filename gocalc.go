package main

// this still doesn't work

import "fmt"

func main() {
	var firstnum float64
	var secondnum float64
	var operator string

	fmt.Println("Please enter your first number")
	fmt.Scan(&firstnum)

	fmt.Println("Please enter your second number")
	fmt.Scan(&secondnum)

	fmt.Println("Please enter your operator")
	fmt.Scan(&operator)

	resulttimes := times(firstnum, secondnum)
	resultdivide := divide(firstnum, secondnum)
	resultadd := add(firstnum, secondnum)
	resultsubtract := subtract(firstnum, secondnum)

	switch operator {
	case "*":
		fmt.Println(resulttimes)
	case "/":
		fmt.Println(resultdivide)
	case "+":
		fmt.Println(resultadd)
	case "-":
		fmt.Println(resultsubtract)
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
