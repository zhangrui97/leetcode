func luckyNumbers (matrix [][]int) []int {
  minr := make([]int, len(matrix))
  for i:= 0; i < len(matrix); i++ { minr[i] = 1e5 }
  maxc := make([]int, len(matrix[0]))
  for i, r := range matrix {
    for j, v := range r {
      if v < minr[i] {
        minr[i] = v
      }
      if v > maxc[j] {
        maxc[j] = v
      }
    }
  }
  var min, max []int
  if len(minr) > len(maxc) {
    min, max = maxc, minr
  } else {
    min, max = minr, maxc
  }
  ans := make([]int, 0, len(min))
  for _, mn := range min {
    for _, mx := range max {
      if mn == mx {
        ans = append(ans, mn)
        break
      }
    }
  }
  return ans
}