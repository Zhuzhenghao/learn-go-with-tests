package main

// Sum cc
func Sum(numbers []int) (sum int) {
	for _, value := range numbers {
		sum += value
	}
	return
}

// SumAll xx
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums

}
