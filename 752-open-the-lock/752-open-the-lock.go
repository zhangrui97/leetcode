func up(str string, pos int) string {
  chs := []rune(str)
  if chs[pos] >= '9' {
    chs[pos] = '0'
  } else {
    chs[pos]++
  }
  return string(chs)
}

func down(str string, pos int) string {
  chs := []rune(str)
  if chs[pos] <= '0' {
    chs[pos] = '9'
  } else {
    chs[pos]--
  }
  return string(chs)
}

func openLock(deadends []string, target string) int {
  dead := make(map[string]bool)
  for _, v := range deadends {
    dead[v] = true
  }
  visited := make(map[string]bool)
  next := make([]string, 0, 10000)
  next = append(next, "0000")
  visited["0000"] = true
  count := 0
  for l := len(next); l > 0; l = len(next) {
    for i:=0; i < l; i++ {
      str := next[0]
      next = next[1:]
      if str == target { return count }
      if _, ok := dead[str]; ok {
        continue
      }      
      for j := 0; j < 4; j ++ {
        upStr := up(str, j)
        if _, ok := visited[upStr]; !ok {
          next = append(next, upStr)
          visited[upStr] = true
        }
        dnStr := down(str, j)
        if _, ok := visited[dnStr]; !ok {
          next = append(next, dnStr)
          visited[dnStr] = true
        }
      }
    }
    count++
  }
  return -1
}