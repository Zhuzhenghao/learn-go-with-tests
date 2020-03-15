package main

// Sum cc
func Sum(numbers []int) (sum int) {
	for _, value := range numbers {
		sum += value
	}
	return
}

// SumAllTails xx
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums

}
