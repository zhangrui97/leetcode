func min(a, b int) int {
  if a > b {
    return b
  } else {
    return a
  }
}
func max(a, b int) int {
  if a < b {
    return b
  } else {
    return a
  }
}
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
  if len(firstList) == 0 || len(secondList) == 0 {
    return [][]int{}
  }  
  ans := make([][]int, 0)
  snd := 0
  for _, v := range firstList {
    s, e := v[0], v[1]
    for ; snd < len(secondList) && secondList[snd][1] < s; snd++ {}
    if snd >= len(secondList) { break }
    fmt.Println(s, e, snd)
    for ; snd < len(secondList) && secondList[snd][0] <= e; {
      ans = append(ans, []int{max(s, secondList[snd][0]), min(e, secondList[snd][1])})
      if secondList[snd][1] < e {
        snd++
      } else {
        break
      }
    }
  }
  return ans
}