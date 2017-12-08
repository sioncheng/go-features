package main

import "fmt"

func main()  {
	ia := [3]int{1,2,3}
	fmt.Println(ia)
	ib := ia
	fmt.Println(ib)
	ia[0] = 10
	fmt.Println(ia)
	fmt.Println(ib)
}