func sum(row []int) int {
  ans := 0
  for _, v := range row {
    ans += v
  }
  return ans
}

func kWeakestRows(mat [][]int, k int) []int {
  m := len(mat)
  hash := make([]int, 0, m)
  for i, r := range mat {
    hash = append(hash, sum(r) * 100 + i)
  }
  sort.Ints(hash)
  for i := 0; i < k; i ++ {
    hash[i] %= 100
  }
  return hash[:k]
}