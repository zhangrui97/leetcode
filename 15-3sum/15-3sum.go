func threeSum(nums []int) [][]int {
  length := len(nums)
  if length < 3 { return [][]int{} }
  sort.Ints(nums)
  ans := make([][]int, 0, length - 2)
  ni := nums[0] - 1
  for i := 0; i < length - 2; i++ {
    for i < length - 2 && nums[i] == ni { i++ }
    ni = nums[i]
    if ni > 0 { break }
    l, r := i+1, length - 1
    for l < r {
      nj := nums[l]
      if (ni + (nj<<1)) > 0 { break }
      nk := nums[r]
      sum := ni + nj + nk
      if sum == 0 { 
        ans = append(ans, []int{ni, nj, nk})
        l ++
        for ; l < r && nums[l] == nj; l++ {}
        r --
        for ; l < r && nums[r] == nk; r-- {}
      } else if sum > 0 {
        r --
        for ; l < r && nums[r] == nk; r-- {}
      } else {
        l ++
        for ; l < r && nums[l] == nj; l++ {}
      }
    }
  }
  return ans
}