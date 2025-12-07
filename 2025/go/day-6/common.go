package day6

func addition(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func multiplication(nums []int) int {
	total := 1
	for _, num := range nums {
		total *= num
	}
	return total
}
