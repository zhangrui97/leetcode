func min(rect []int) int {
  if rect[0] < rect[1] {
    return rect[0]
  } else {
    return rect[1]
  }
}

func countGoodRectangles(rectangles [][]int) int {
  sq := make([]int, 0, len(rectangles))
  largest := 0
  for _, v := range rectangles {
    l := min(v)
    sq = append(sq, l)
    if l > largest {
      largest = l
    }
  }
  ans := 0
  for _, v := range sq {
    if v == largest {
      ans ++
    }
  }
  return ans
}