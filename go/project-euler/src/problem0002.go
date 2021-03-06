package main

import "fmt"

/*
Each new term in the Fibonacci sequence is generated by adding the previous two terms.
By starting with 1 and 2, the first 10 terms will be:

1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

By considering the terms in the Fibonacci sequence whose values do not exceed four million,
find the sum of the even-valued terms.
*/

func main2() {
	var sum = 10
	var n3 = 2
	var n6 = 8

	for ; n6 < 1000000; {
		var tmp = n3 + n6 * 4
		sum += tmp
		n3 = n6
		n6 = tmp
	}

	fmt.Printf("Result: %d", sum)
}
