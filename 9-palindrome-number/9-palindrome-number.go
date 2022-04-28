func isPalindrome(x int) bool {
  if x < 0 { return false }
  if x < 10 { return true }
  rev := 0
  for rem := x; rem > 0; rem /= 10 {
    rev = rev * 10 + rem % 10
  }
  return rev == x
}