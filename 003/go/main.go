package main

import "fmt"

func main() {
	primes := getPrimeFactors([]int64{600851475143})
	largest := max(primes)
	fmt.Printf("all prime factors: %v\n", primes)
	fmt.Printf("largest prime factor: %v\n", largest)
}

// getPrimeFactors is a function that will get all the prime factors of a
// number and return them in a list of int64 values.
func getPrimeFactors(numbers []int64) []int64 {

	// set up an empty list that will hold all the divisors of a number
	var divisors []int64

	// loop through the values in the "number" list, ignore index
	for _, num := range numbers {
		// start with a divisor of 2, which is the
		// smallest prime number
		var div int64 = 2

		// loop using the value of "div" to see if the number "num"
		// from the list of numbers can be evenly divided by div. if
		for ; div < num; div++ {
			// add div and the quotient to the list of divisors
			if num%div == 0 && num != div {
				divisors = insertUnique(
					divisors, div, (num / div))
			}

			// if the divisor is larger than its quotient, then
			// stop checking for this value of div
			if div > (num / div) {
				break
			}
		}
	}

	// if the length of "divisors" is 0, then that means that all the
	// values in "numbers" can only be divided by 1 and themselves,
	// ie: they are all prime numbers
	if len(divisors) == 0 {
		return numbers
	} else {
		// recursive function call to get more primes
		return getPrimeFactors(divisors)
	}
}

// containsVal checks it a value is already present in a list and returns
// the index of a value in that list, otherwise returns -1 if not found
func containsVal(nums []int64, val int64) int {
	ret := -1

	for i, v := range nums {
		if v == val {
			ret = i
			break
		}
	}

	return ret
}

// insertUnique inserts a value into a list if that value is not already
// present in the list
func insertUnique(nums []int64, vals ...int64) []int64 {
	for _, v := range vals {
		if containsVal(nums, v) == -1 {
			nums = append(nums, v)
		}
	}

	return nums
}

// max finds the largest value in a list
func max(numbers []int64) int64 {
	var max int64

	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	return max
}

/*

Problem 3 - Largest prime factor

The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?

*/
