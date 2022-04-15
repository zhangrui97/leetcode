func abs(a int) int {
  if a < 0 {
    return -a
  } else {
    return a
  }
}

func threeSumClosest(nums []int, target int) int {
  sort.Ints(nums)
  length := len(nums)
  ans := 3000
  for i := 0; i < length; i++ {
    ai := nums[i]
    l, r := i+1, length-1
    for l < r {
      sum := ai + nums[l] + nums[r]
      if abs(sum - target) < abs(ans - target) {
        ans = sum
      }
      if sum > target {
        r--
      } else {
        l++
      }
    }
  }
  return ans
}