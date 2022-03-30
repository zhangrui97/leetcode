func countPairs(nums []int, k int) int {
  idxMap := make(map[int][]int)
  for i, v := range nums {
    idxMap[v] = append(idxMap[v], i)
  }
  ans := 0
  for _, v := range idxMap {
    for i := 0; i < len(v) - 1; i++ {
      idxi := v[i]
      for j := i+1; j < len(v); j++ {
        if idxi * v[j] % k == 0 {
          ans++
        }
      }
    }
  }
  return ans
}