func numTrees(n int) int {
  switch n {
    case 0: return 1
    case 1: return 1
    case 2: return 2
    case 3: return 5
    default:
      ans := 0
      for i := 0; i < n; i++ {
        ans += numTrees(i) * numTrees(n-1-i)
      }
      return ans
  }
}