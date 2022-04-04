func canPlaceFlowers(flowerbed []int, n int) bool {
  if n == 0 { return true }
  l := len(flowerbed)
  if l == 1 {
    return flowerbed[0] == 0
  }  
  if flowerbed[0] == 0 && flowerbed[1] == 0 {
    flowerbed[0] = 1
    n--
  }
  if n== 0 { return true }
  if flowerbed[l-1] == 0 && flowerbed[l-2] == 0 {
    flowerbed[l-1] = 1
    n--
  }
  if n== 0 { return true }
  for i := 1; i < l-1; i++ {
    if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
      n--
      if n== 0 { return true }
      i++
    }
  }
  return n == 0
}