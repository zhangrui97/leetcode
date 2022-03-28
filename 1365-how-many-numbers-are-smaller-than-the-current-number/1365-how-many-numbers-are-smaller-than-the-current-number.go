func indexOf(sorted []int, value int) int {
  left, right := 0, len(sorted) -1
  for (left <= right) {
    mid := left + (right - left) / 2
    if sorted[mid] < value {
      left = mid + 1
    } else {
      right = mid - 1
    }
  }
  return left
}

func smallerNumbersThanCurrent(nums []int) []int {
  sorted := make([]int, len(nums))
  copy(sorted, nums)
  sort.Ints(sorted)
  ans := make([]int, len(nums))
  for i, v := range nums {
    ans[i] = indexOf(sorted, v)  
  }
  return ans
}