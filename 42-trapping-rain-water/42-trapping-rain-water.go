func trap(height []int) int {
  length := len(height)
  if length < 2 { return 0 }
  sum := 0
  l, r := 0, length - 1
  lmax := height[l]
  rmax := height[r]
  for l <= r {
    lv, rv := height[l], height[r]
    if lv >= lmax {
      lmax = lv
      l++
    } else if lmax <= rmax {
      sum += lmax - height[l]
      l ++
    }
    if l > r { break }
    if rv >= rmax {
      rmax = rv
      r--
    } else if rmax <= lmax {
      sum += rmax - height[r]
      r --
    }
  }
  return sum
}