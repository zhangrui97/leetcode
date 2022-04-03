func thirdMax(nums []int) int {
  max, max1, max2 := math.MinInt32-1, math.MinInt32-1, math.MinInt32-1
  for _, v := range nums {
    if v > max {
      max, max1, max2 = v, max, max1
    } else if v == max {
      continue
    } else if v > max1 {
      max1, max2 = v, max1
    } else if v == max1 {
      continue
    } else if v >= max2 {
      max2 = v
    }
  }
  if max2 >= math.MinInt32 {
    return max2
  } else {
    return max
  }
}