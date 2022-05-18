func numTrees(n int) int {
  cach := make(map[int]int, n+1)
  cach[0] = 1
  var calc func(m int)int
  calc = func(m int)int {
    if v, ok := cach[m]; ok {
      return v
    }
    ans := 0
    for i := 0; i < m; i++ {
      ans += calc(i) * calc(m-1-i)
    }
    cach[m] = ans
    return ans
  }
  return calc(n)
}