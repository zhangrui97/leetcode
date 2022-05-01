func isValid(s string) bool {
  stack := make([]rune, 0, len(s))
  for _, ch := range s {
    switch ch {
      case '(': stack = append(stack, ')')
      case '[': stack = append(stack, ']')
      case '{': stack = append(stack, '}')
      default:
        if len(stack) == 0 || ch != stack[len(stack)-1] {
          return false
        } else {
          stack = stack[:len(stack)-1]
        }
    }
  }
  return len(stack) == 0
}