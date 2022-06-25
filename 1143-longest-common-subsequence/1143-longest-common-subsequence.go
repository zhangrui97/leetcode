func max(a, b int)int {
  if a > b {
    return a
  } else {
    return b
  }
}

func longestCommonSubsequence(text1 string, text2 string) int {
  n1, n2 := len(text1), len(text2)
  memo := make(map[[2]int]int)
  var dp func(i, j int)int
  dp = func(i, j int)int {
    index := [2]int{i, j}
    if v, ok := memo[index]; ok { return v }
    if i < 0 || j < 0 { return 0 }
    if text1[i] == text2[j] { 
      v := dp(i-1, j-1) + 1
      memo[index] = v
      return v }
    v := max(dp(i-1, j), dp(i, j-1))
    memo[index] = v
    return v
  }
  return dp(n1-1, n2-1)
}