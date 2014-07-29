/*
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

package main

import "fmt"

func main() {
	const n int64 = 600851475143
	var newnum int64 = n
	var largestFact int64 = 0

	var counter int64 = 2

	for counter*counter <= newnum {
		if newnum%counter == 0 {
			newnum = newnum / counter
			largestFact = counter

		} else {
			counter++
		}
	}

	if newnum > largestFact {
		largestFact = newnum
	}

	fmt.Println("Largest factor: ", largestFact)
}
