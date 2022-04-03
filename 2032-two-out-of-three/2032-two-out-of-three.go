func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
  flags := [101]int{}
  for _, v := range nums1 {
    flags[v] = 1
  }
  for _, v := range nums2 {
    flags[v] |= 2
  }
  for _, v := range nums3 {
    flags[v] |= 4
  }
  result := make([]int, 0, 100)
  for i, v := range flags {
    switch v {
    case 3:
      result = append(result, i)  
    case 6:
      result = append(result, i)  
    case 5:
      result = append(result, i)  
    case 7:
      result = append(result, i)  
    }
  }
  return result
}