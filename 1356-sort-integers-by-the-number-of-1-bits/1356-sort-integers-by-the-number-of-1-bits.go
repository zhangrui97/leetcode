import "math/bits"

func sortByBits(arr []int) []int {
  sort.Slice(arr, func(i, j int) bool {
    return bits.OnesCount(uint(arr[i])) << 16 | arr[i] < bits.OnesCount(uint(arr[j])) << 16 | arr[j]})    
  return arr
}