func trap(height []int) int {
  length := len(height)
  if length < 2 { return 0 }
  maxStack := make([]int, 0, length)
  sum := 0
  for i, v := range height {
    if i == 0 || v < height[maxStack[len(maxStack) - 1]] {
      maxStack = append(maxStack, i)
    } else {
      for j, max := range maxStack {
        if v >= height[max] {
          maxStack[j] = i
          maxStack = maxStack[0:j+1]
          break
        }
      }
    }
    sum += height[maxStack[0]] - v
  }
  for i := len(maxStack) - 1; i > 0; i-- {
    sum -= (maxStack[i] - maxStack[i-1])*(height[maxStack[0]] - height[maxStack[i]])
  }
  return sum
}