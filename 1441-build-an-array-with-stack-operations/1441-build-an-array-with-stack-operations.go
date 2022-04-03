func buildArray(target []int, n int) []string {
  cur := 1
  ans := make([]string, 0, 2 * n)
  for _, v := range target {
    for ; cur < v; cur++ {
      ans = append(ans, "Push")
      ans = append(ans, "Pop")
    }
    ans = append(ans, "Push")
    cur++
  }
  return ans
}