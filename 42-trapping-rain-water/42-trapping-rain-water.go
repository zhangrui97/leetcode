func trap(height []int) int {
  length := len(height)
  sum := 0
  l, lmax, rmax, r := 0, 0, 0, length - 1
  for l < r {
    lv, rv := height[l], height[r]
    if lv < rv {
      if lv > lmax {
        lmax = lv
      } else {
        sum += lmax - lv
      }
      l++
    } else {
      if rv >= rmax {
        rmax = rv
      } else {
        sum += rmax - rv
      }
      r --
    }
  }
  return sum
}