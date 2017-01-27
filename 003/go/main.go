package main

import "fmt"

func main() {
	fmt.Printf("answer: %v\n", getPrimeFactors([]int64 {600851475143}))
}

func getPrimeFactors(numbers []int64) []int64 {

	var divisors []int64
	for _, num := range numbers {
		var div int64 = 2

		for ;div < num; div++ {
			if num%div == 0 && num != div {
				if containsVal(divisors, div) == -1 {
					fmt.Printf("adding %v\n", div)
					divisors = append(divisors, div)
				}
			}
		}
	}
	
	if (len(divisors) == 0) {
		return numbers
	} else {
		fmt.Println(divisors)
		return getPrimeFactors(divisors[len(divisors)-1:])
	}
}


// containsVal : returns the index of a value in the slice, 
// otherwise returns -1 if not found
func containsVal(nums []int64, val int64) int {
	ret := -1

	for i, v := range nums {
		if (v == val) {
			ret = i
			break
		}
	}

	return ret
}
