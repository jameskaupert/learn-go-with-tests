package arraysandslices

func Sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, s := range numbersToSum {
		if len(s) == 0 {
			sums = append(sums, 0)
		} else {
			tail := s[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
