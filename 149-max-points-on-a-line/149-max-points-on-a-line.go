type Line struct {
  k float64
  b float64
}

func maxPoints(points [][]int) int {
  l := len(points)
  if l <= 2 { return l }
  res := make(map[Line][]int)  
  counter := 2
  skip := make(map[[2]int]bool)
  for i := 0; i < l-1; i++ {
    xi, yi := points[i][0], points[i][1]
    for j := i+1; j < l; j++ {
      if _, ok := skip[[2]int{i, j}]; ok { continue }
      xj, yj := points[j][0], points[j][1]
      line := Line{math.Inf(1), float64(xi)}
      if xi != xj {
        line.k = float64(yj-yi)/float64(xj-xi)
        line.b = float64(yi) - float64(xi)*line.k
      }
      if v, ok := res[line]; ok {
        ll := len(v)
        for k:= 1; k < ll; k++ {
          skip[[2]int{v[k], j}] = true
        }
        res[line] = append(v, j)
        if ll >= counter {
          counter = ll + 1
        }
      } else {
        res[line] = []int{i, j}
      }
    }
  }
  return counter
}