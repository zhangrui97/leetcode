func min(a, b int)int {
  if a < b {
    return a
  } else {
    return b
  }
}

func minimumDeleteSum(s1 string, s2 string) int {
  n := len(s2)
  dpPre, dpCur := make([]int, n+1), make([]int, n+1) 
  sum := 0
  for i, ch := range s2 {
    sum += int(ch)
    dpCur[i+1] = sum
  }
  sum = 0
  for _, ch1 := range s1 {
    dpPre, dpCur = dpCur, dpPre
    sum += int(ch1)
    dpCur[0] = sum
    for i2, ch2 := range s2 {
      if ch1 == ch2 {
        dpCur[i2+1] = dpPre[i2]
      } else {
        dpCur[i2+1] = min(dpCur[i2] + int(ch2), dpPre[i2+1] + int(ch1))
      }
    }
  }
  return dpCur[n]
}