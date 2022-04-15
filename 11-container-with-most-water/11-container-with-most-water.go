func min(a, b int) int {
  if a > b {
    return b
  } else {
    return a
  }
}

func maxArea(height []int) int {
  l, r := 0, len(height) - 1
  h := min(height[l], height[r])
  ans := (r - l)  * h
  for l < r {
    for ; l < r && height[l] <= h; l++ {}
    for ; l < r && height[r] <= h; r-- {}
    if l >= r { break }
    h = min(height[l], height[r])
    area := (r-l) * h
    if area > ans {
      ans = area
    }
  }
  return ans
}