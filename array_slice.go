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

	sa := make([]int, 3)
	sa[0] = 1
	sa[1] = 2
	sa[2] = 3
	fmt.Println(sa)
	sb := append(sa, 4)
	fmt.Println(sb)
	fmt.Println(sa)
	sa[0] = 10
	fmt.Println(sa)
	fmt.Println(sb)
	sc := sa
	sa[0] = 20
	fmt.Println(sa)
	fmt.Println(sc)
}