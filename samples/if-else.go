package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if 8%4 == 0 {
		fmt.Println("divisible")
	}

	if num := 9; num < 9 {
		fmt.Println(num, " negative")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, " multiple digits")
	}
}
