func smallestSubsequence(s string) string {
  var counter [127]int
  for _, v := range s {
    counter[v]++
  }
  stack := make([]rune, 0, len(s))
  var inStack [127]bool
  for _, v := range s {
    counter[v]--
    if inStack[v] { continue }
    for l := len(stack); l > 0 && counter[stack[l-1]] > 0 && v < stack[l-1]; l-- {
      inStack[stack[l-1]] = false
      stack = stack[:l-1]
    }
    stack = append(stack, v)
    inStack[v] = true
  }
  return string(stack)
}