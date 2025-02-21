package main

import "fmt"

func main() {
	var length int
	var width int
	var i int
	var LastSymb string = " "
	var NumString int = 1
	var OutputLine string

	fmt.Printf("Enter length: ")
	fmt.Scanln(&length)
	fmt.Printf("Enter width: ")
	fmt.Scanln(&width)

	for i = 1; i <= width*length; i++ {
		switch LastSymb {
		case " ":
			OutputLine = OutputLine + "#"
			LastSymb = "#"
		case "#":
			OutputLine = OutputLine + " "
			LastSymb = " "
		}
		if i%length == 0 {
			OutputLine = OutputLine + "\n"
			NumString++
			if length%2 == 0 && NumString%2 == 0 {
				LastSymb = "#"
			} else if length%2 == 0 && NumString%2 != 0 {
				LastSymb = " "
			}
		}
	}
	fmt.Println(OutputLine)
}
