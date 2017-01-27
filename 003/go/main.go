package main

import "fmt"

func main() {
	ans := max(getPrimeFactors([]int64 {600851475143}))
	fmt.Printf("answer: %v\n", ans)
}

func getPrimeFactors(numbers []int64) []int64 {

	var divisors []int64
	for _, num := range numbers {
		var div int64 = 2

		for ;div < num; div++ {
			if num%div == 0 && num != div {
				divisors = insert(divisors, div, (num / div))
			}
			
			if (div > (num / div)) {
				break
			}
		}
	}
	
	if (len(divisors) == 0) {
		return numbers
	} else {
		return getPrimeFactors(divisors)
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


func insert(nums []int64, vals ...int64) []int64{
	for _, v := range vals {
		if (containsVal(nums, v) == -1) {
			nums = append(nums, v)
		}
	}
	
	return nums
}

func max(numbers []int64) int64 {
	var max int64
	
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	
	return max
}
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		
		