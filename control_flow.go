package main

import "fmt"

func main()  {
	a := 1
	if a == 1 {
		fmt.Println("a == 1")
	} else {
		fmt.Println("a <> 1")
	}
	for i:= 0 ; i < 3; i ++ {
		fmt.Print(i)
		fmt.Print(" ")
	}
	fmt.Println("")

	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	}
}