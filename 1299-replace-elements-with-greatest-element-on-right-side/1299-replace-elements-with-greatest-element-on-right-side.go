func replaceElements(arr []int) []int {
  n := len(arr)
  max := arr[n-1]
  arr[n-1] = -1
  for i := n - 2; i >= 0; i-- {
    v := arr[i]
    arr[i] = max
    if v > max {
      max = v
    }
  }
  return arr
}