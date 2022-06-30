func canPartition(nums []int) bool {
  sum, max := 0, 0
	for _, num := range nums {
		sum += num
    if num > max { max = num }
	}
	if sum%2 != 0 || max+max > sum {
		return false
	}
	sum = sum / 2
  fmt.Println(sum)
	for range nums {
    ans := 0
    fmt.Println(nums)
		for _, num := range nums {
			if num+ans <= sum {
				ans += num
			}
			if ans == sum {
				return true
			}
		}
		nums = append(nums, nums[0])
		nums = nums[1:]
	}

	return false
}