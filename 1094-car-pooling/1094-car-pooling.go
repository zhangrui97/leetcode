func carPooling(trips [][]int, capacity int) bool {
  diff := make([]int, 1001)
  for _, v := range trips {
    diff[v[1]] += v[0]
    if v[2] <= 1000 {
      diff[v[2]] -= v[0]
    }
  }
  sum := 0
  for _, v := range diff {
    sum += v
    if sum > capacity {
      return false
    }
  }
  return true
}