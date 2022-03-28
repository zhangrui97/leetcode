func mostWordsFound(sentences []string) int {
  ans := 0
  for _, v := range sentences {
    l := len(strings.Split(v, " "))
    if l > ans {
      ans = l
    }
  }
  return ans
}