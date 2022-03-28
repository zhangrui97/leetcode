func finalValueAfterOperations(operations []string) int {
  ans := 0
  for _, v := range operations {
    if v[1] == '+' {
      ans += 1
    } else {
      ans -= 1
    }
  }
  return ans
}