package main

import (
	"fmt"
	"math"
)

// fmt anf math are libraries of main package where we have methods
// main package is not library its special or executable package go build main->gives exe file
//
//	hey this just  basics of go
func main() {
	var x int
	var y string
	// taking input in golang using fmt library from main package and its methods like scanf etc
	_, err := fmt.Scanf("%d %s", &x, &y)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println("You entered x:", x, " and y:", y)

	var g string = "hello world"
	a := "10"
	fmt.Println(a, g)
	var f float64 = math.Sqrt(625)
	fmt.Println("the square root of 625 is ", f)
	// constants in golang
	const Pi = 3.14
	if f < 0 {
		fmt.Println("f is neg")
	} else if f == 0 {
		fmt.Println("f is 0")
	} else {
		fmt.Println("f is pos")
	}
	res := add(10, 200)
	fmt.Println(res)
	show := func(a, b int) {
		fmt.Println(a + b)
	}
	show(10, 20)
	// arrays declaration
	var str [5]string
	str[0] = "adithya"
	str[1] = "doreamon"
	fmt.Println(str)
	//  slice-> dynamic arraay
	primes := make([]int, 5)
	for i := 0; i < len(primes); i++ {
		_, err = fmt.Scanf("%d", &primes[i])
		// or _,err=fmt.Scan(&primes[i])
		if err != nil {
			fmt.Println("error in reading input")
			return
		}
	}
	fmt.Println("The prime numbers are:", primes)
	// maps in golang is like a hashmap
	// m := make(map[string]int)
	freq := make(map[int]int)
	// counting freq of primes array
	for i := 0; i < len(primes); i++ {
		freq[primes[i]]++
	}
	fmt.Println(freq)
	//    structs in goLang
	type cordinates struct {
		x int
		y int
	}
	p1 := cordinates{2, 4}
	fmt.Println(p1)

}
func add(x int, y int) int {
	return (x + y)
}
