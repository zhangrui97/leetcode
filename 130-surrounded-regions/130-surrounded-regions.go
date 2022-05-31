func solve(board [][]byte)  {
  m := len(board)
  n := len(board[0])
  if m <= 2 || n <= 2 {
    return
  }
  ans := make(map[int]bool)
  next := make([]int, 0, (m+n+2)<<1)
  for j := 0; j < n; j++ {
    if board[0][j] == 'O' {
      next = append(next, (1<<16|j))
      ans[j] = true
    }
    if board[m-1][j] == 'O' {
      next = append(next, ((m-2)<<16)|j)
      ans[((m-1)<<16)|j] = true
    }
  }
  for i := 1; i < m-1; i++ {
    if board[i][0] == 'O' {
      next = append(next, i<<16|1)
      ans[i<<16] = true
    }
    if board[i][n-1] == 'O' {
      next = append(next, (i<<16)|(n-2))
      ans[(i<<16)|(n-1)] = true
    }
  }
  for l:=len(next); l > 0; l=len(next) {
    for k := 0; k < l; k++ {
      v := next[0]
      next = next[1:]
      if ans[v] { continue }
      i, j := (v & 0xFFFF0000) >> 16, v & 0x0000FFFF
      if board[i][j] != 'O' { continue }
      next = append(next, ((i-1)<<16)|j)
      next = append(next, ((i+1)<<16)|j)
      next = append(next, (i<<16)|(j-1))
      next = append(next, (i<<16)|(j+1))
      ans[(i<<16)|j] = true
    }
  }
  for i, r := range board {
    for j, _ := range r {
      if _, ok := ans[(i << 16) | j]; ok {
        r[j] = 'O'
      } else {
        r[j] = 'X'
      }
    }
  }
}