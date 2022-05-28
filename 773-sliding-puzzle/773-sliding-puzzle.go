func slidingPuzzle(board [][]int) int {
  var target int
  for i := 0; i < 2; i++ {
    for j, v := range board[i] {
      target |= v << (20 - 12*i - 4*j)
      if v == 0 {
        target |= (3*i+j) << 24
      }
    }
  }
  visited := make(map[int]bool)
  next := []int{0x05123450} 
  count := 0
  for l := len(next); l > 0; l = len(next) {
    for i := 0; i < l; i++ {
      v := next[0]
      next = next[1:]
      if v == target { return count }
      switch (v & 0xFF000000) {
        case 0x00000000:
          n := ((v & 0x000F0000) << 4 | v) & 0x00F0FFFF | 0x01000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
          n = ((v & 0x00000F00) << 12 | v) & 0x00FFF0FF | 0x03000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
        case 0x01000000:
          n := ((v & 0x00F00000) >> 4 | v) & 0x000FFFFF
          if _, ok := visited[n]; !ok { next = append(next, n) }
          n = ((v & 0x0000F000) << 4 | v) & 0x00FF0FFF | 0x02000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
          n = ((v & 0x000000F0) << 12 | v) & 0x00FFFF0F | 0x04000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
        case 0x02000000:
          n := ((v & 0x000F0000) >> 4 | v) & 0x00F0FFFF | 0x01000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
          n = ((v & 0x0000000F) << 12 | v) & 0x00FFFFF0 | 0x05000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
        case 0x03000000:
          n := ((v & 0x00F00000) >> 12 | v) & 0x000FFFFF
          if _, ok := visited[n]; !ok { next = append(next, n) }
          n = ((v & 0x000000F0) << 4 | v) & 0x00FFFF0F | 0x04000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
        case 0x04000000:
          n := ((v & 0x000F0000) >> 12 | v) & 0x00F0FFFF | 0x01000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
          n = ((v & 0x00000F00) >> 4 | v) & 0x00FFF0FF | 0x03000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
          n = ((v & 0x0000000F) << 4 | v) & 0x00FFFFF0 | 0x05000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
        case 0x05000000:
          n := ((v & 0x0000F000) >> 12 | v) & 0x00FF0FFF | 0x02000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
          n = ((v & 0x000000F0) >> 4 | v) & 0x00FFFF0F | 0x04000000
          if _, ok := visited[n]; !ok { next = append(next, n) }
      }
      visited[v] = true
    }
    count++
  }
  return -1
}