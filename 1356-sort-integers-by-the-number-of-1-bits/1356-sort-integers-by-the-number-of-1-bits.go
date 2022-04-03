func bits(num int) int {
  if num == 0 {
    return 0
  }
  return 1 + bits(num & (num-1))
}

func sortByBits(arr []int) []int {
  sort.Slice(arr, func(i, j int) bool {
    return bits(arr[i]) << 16 | arr[i] < bits(arr[j]) << 16 | arr[j]
  })    
  return arr
}