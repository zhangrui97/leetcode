func countConsistentStrings(allowed string, words []string) int {
  count := 0
  for _, word := range words {
    count ++
    for _, r := range word {
      if !strings.ContainsRune(allowed, r) {
        count --
        break
      }      
    }
  } 
  return count
}