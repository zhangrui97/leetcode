func totalNQueens(n int) int {
  ans := 0
  res := make([][]byte, n)
  for i := 0; i < n; i++ {
    res[i] = make([]byte, n)
    for j:= 0; j < n; j++ {
      res[i][j] = '.'
    }
  }
  conflict := func(i, j int)bool {
    for l := 0; l < i; l ++ {
      if res[l][j] == 'Q' { return true }
      dist := i - l
      if j - dist >= 0 && res[l][j-dist] == 'Q' { return true }
      if j + dist < n && res[l][j+dist] == 'Q' { return true }
    }
    for k := 0; k < n; k++ {
      if res[i][k] == 'Q' { return true }
    }
    for l := i+1; l < n; l++ {
      dist := l - i
      if j - dist >= 0 && res[l][j-dist] == 'Q' { return true }
      if j + dist < n && res[l][j+dist] == 'Q' { return true }
    }
    return false
  }
  var backtrack func(i int)
  backtrack = func(i int) {
    if i == n {
      ans++
      return
    }
    for j := 0; j < n; j++ {
      if conflict(i, j) { continue }
      res[i][j] = 'Q'
      backtrack(i + 1)
      res[i][j] = '.'
    }
  }
  backtrack(0)
  return ans
}