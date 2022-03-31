func targetIndices(nums []int, target int) []int {
  sort.Ints(nums)
  l, r := 0, len(nums) - 1
  for l <= r {
    mid := (l + r) / 2
    if nums[mid] < target {
      l = mid + 1
    } else {
      r = mid - 1
    }
  }
  left := l
  l, r = 0, len(nums) - 1
  for l <= r {
    mid := (l+r) / 2
    if nums[mid] > target {
      r = mid - 1
    } else {
      l = mid + 1
    }
  }
  if left > r {
    return []int{}
  } else {
    ans := make([]int, 0, r + 1 - left)
    for i := left; i <= r; i++ {
      ans = append(ans, i)
    }
    return ans
  }
}