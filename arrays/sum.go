package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(allNumbers ...[]int) []int {
	var sums []int

	for _, numbers := range allNumbers {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(allNumbers ...[]int) []int {
	var sums []int

	for _, numbers := range allNumbers {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}

	return sums
}
