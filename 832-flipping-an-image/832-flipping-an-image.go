func flipAndInvertImage(image [][]int) [][]int {
  n := len(image)
  isOdd := (n % 2 == 1)
  for _, v := range image {
    for i := 0; i < n/2; i++ {
      v[i], v[n-1-i] = 1 - v[n-1-i], 1 - v[i]
    }
    if isOdd {
      v[n/2] = 1 - v[n/2]
    }
  }
  return image
}