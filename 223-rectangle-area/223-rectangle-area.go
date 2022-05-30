func min(a, b int) int {
  if a > b { 
    return b 
  } else {
    return a
  }
}
func max(a, b int) int {
  if a > b {
    return a
  } else {
    return b
  }
}

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
  aa := (ax2-ax1)*(ay2-ay1)
  ab := (bx2-bx1)*(by2-by1)
  intersect := 0
  if bx1 < ax2 && bx2 > ax1 && by2 > ay1 && by1 < ay2 {
    intersect = (min(ax2, bx2) - max(ax1, bx1))*(min(ay2, by2) - max(ay1, by1))
  }
  return aa + ab - intersect
}