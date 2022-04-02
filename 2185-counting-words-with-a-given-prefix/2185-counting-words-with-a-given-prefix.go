func prefixCount(words []string, pref string) int {
  ans := 0
  for _, v := range words {
    if strings.HasPrefix(v, pref) {
      ans ++
    }
  }  
  return ans
}