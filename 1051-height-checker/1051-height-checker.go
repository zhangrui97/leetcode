func heightChecker(heights []int) int {
  expected := make([]int, len(heights))
  copy(expected, heights)
  sort.Ints(expected)
  ans := 0
  for i := 0; i < len(heights); i++ {
    if expected[i] != heights[i] {
      ans ++
    }
  }
  return ans
}