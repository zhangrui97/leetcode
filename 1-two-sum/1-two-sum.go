func twoSum(nums []int, target int) []int {
  cache := make(map[int]int)
  for i, v := range nums {
    if j, ok := cache[v]; ok {
      return []int{j, i}
    }
    cache[target - v] = i
  }
  return []int{}
}