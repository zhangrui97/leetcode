func min(a, b int)int {
  if a > b {
    return b
  } else {
    return a
  }
}

func minDistance(word1 string, word2 string) int {
  n := len(word2)
  dpPre, dpCur := make([]int, n + 1), make([]int, n + 1)
  for i := 1;i <= n; i++ { dpCur[i] = i }
  for i, ch1 := range word1 {
    dpPre, dpCur = dpCur, dpPre
    dpCur[0] = i + 1
    for j, ch2 := range word2 {
      if ch1 == ch2 {
        dpCur[j+1] = dpPre[j]
      } else {
        dpCur[j+1] = min(dpPre[j+1], dpCur[j]) + 1
      }
    }
  }
  return dpCur[len(word2)]
}