func permute(nums []int) [][]int {
  n := len(nums)
  if n == 1 { return [][]int{nums} }
  result := make([][]int, 0, (n-1)*n)    
  var backtrack func(i int, path []int)
  backtrack = func(i int, path []int) {
    if i == 0 {
      res := make([]int, n)
      copy(res, path)
      result = append(result, res)
      return
    }
    for j := 0; j <= i; j++ {
      path[j], path[i] = path[i], path[j]
      backtrack(i-1, path)
      path[j], path[i] = path[i], path[j]
    }
  }
  backtrack(n-1, nums)
  return result
}