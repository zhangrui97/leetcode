func absDiff(a, b int) int {
  if a > b {
    return a - b
  } else {
    return b - a
  }
}
func countGoodTriplets(arr []int, a int, b int, c int) int {
  l := len(arr)
  ans := 0
  for i := 0; i < l - 2; i++ {
    vi := arr[i]
    for j := i+1; j < l - 1; j++ {
      vj := arr[j]
      if absDiff(vi, vj) > a {
        continue
      }
      for k:= j+1; k < l; k++ {
        vk := arr[k]
        if absDiff(vj, vk) <= b && absDiff(vi, vk) <= c {
          ans ++
        }
      }
    }
  }
  return ans
}