func countVowelStrings(n int) int {
  twoLetters := (n-1) * 10
  threeLetters := twoLetters*(n-2)/2
  fourLetters := threeLetters*(n-3)/6
  fiveLetters := fourLetters*(n-4)/20
  return 5 + twoLetters + threeLetters + fourLetters + fiveLetters
}