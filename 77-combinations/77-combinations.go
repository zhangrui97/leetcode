func combine(n int, k int) [][]int {
  ans := [][]int{}
  var backtrack func(path int, i, l int)
  backtrack = func(path int, i, l int) {
    if l == k {
      res := make([]int, k)
      for j, k, p := 0, 0, path; p > 0; j, p = j+1, p >> 1 {
        if p & 1 > 0 { 
          res[k] = j+1
          k += 1
        }
      }
      ans = append(ans, res)
      return
    }
    for j := i; j < n; j++ {
      backtrack(path | (1 << j), j+1, l+1)
    }
  }
  backtrack(0, 0, 0)
  return ans
}