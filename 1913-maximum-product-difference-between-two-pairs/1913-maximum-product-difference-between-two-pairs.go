func maxProductDifference(nums []int) int {
  maxIndex, minIndex := 0, 0
  for i, v := range nums {
    if v > nums[maxIndex] {
      maxIndex = i
    }
    if v < nums[minIndex] {
      minIndex = i
    }
  }
  max2Index, min2Index := minIndex, maxIndex
  for i, v := range nums {
    if i != maxIndex && v > nums[max2Index] {
      max2Index = i
    }
    if i != minIndex && v < nums[min2Index] {
      min2Index = i
    }
  }
  
  return nums[maxIndex]*nums[max2Index] - nums[minIndex]*nums[min2Index]
}