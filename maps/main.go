package main

import "fmt"

func printMap(m map[int]int) {
	for k, v := range m {
		fmt.Printf("%v, %v\n", k, v)
	}
}

func main() {
	var one map[int]int      // one has value nil, NOT an initialized empty map
	two := make(map[int]int) // an initialized empty map
	fmt.Printf("%+v, %+v\n", one, two)

	//one[1] = 1 -> panic: assignment to entry in nil map
	two[2] = 2

	three := map[int]int{
		3: 3,
		4: 4,
		5: 5,
		6: 6,
	}
	delete(three, 4)
	fmt.Printf("%+v, %+v, %+v\n", one, two, three)

	printMap(three)
}

/*
all keys in map must have the same data type, so do values
reference type
*/
