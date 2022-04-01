func GCD(a, b int) int {
  if b == 0 {
    return a
  }
  return GCD(b, a%b)
}
func findGCD(nums []int) int {
  min, max := 1000, 1
  for _, v := range nums {
    if v < min {
      min = v
    }
    if v > max {
      max = v
    }
  }
  return GCD(max, min)
}