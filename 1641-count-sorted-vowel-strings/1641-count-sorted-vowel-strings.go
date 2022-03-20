func countVowelStrings(n int) int {
  /*
    put all 5 vowels to the string of length (n + 5), remove all of them once in the string will be result
    so the question becomes find the composition of four positions from the (n + 5 - 1) potential insert positions
    The answer is C(n+4, 4), it is -
    (n+1)*(n+2)*(n+3)*(n+4)/(1*2*3*4) 
      = (n*n + 5*n + 4)(n*n + 5*n +6)/24
      = (((n*n + 5*n)+10)*(n*n + 5*n) + 24)/24
  */
  x := n * (n + 5)
  return (x * (x + 10) + 24) / 24
}