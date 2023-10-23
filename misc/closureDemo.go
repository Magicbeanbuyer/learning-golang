package main

import "fmt"

func makeMult(base int) func(int) int {
	f := func(factor int) int {
		return base * factor
	}
	return f
}

func wrapper_cl() {
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

/*
0 0
2 3
4 6

*/
