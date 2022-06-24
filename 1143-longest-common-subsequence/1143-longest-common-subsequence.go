func max(a, b int)int {
  if a > b {
    return a
  } else {
    return b
  }
}

func longestCommonSubsequence(text1 string, text2 string) int {
  n1, n2 := len(text1), len(text2)
  dpPre, dpCur := make([]int, n2+1), make([]int, n2+1) 
  for i := 1; i <= n1; i++ {
    dpPre, dpCur = dpCur, dpPre
    for j := 1; j <= n2; j++ {
      if text1[i-1] == text2[j-1] {
        dpCur[j] = dpPre[j-1] + 1
      } else {
        dpCur[j] = max(dpPre[j], dpCur[j-1])
      }
    }
  }
  return dpCur[n2]
}