func truncateSentence(s string, k int) string {
  return strings.Join(strings.Fields(s)[:k], " ")  
}