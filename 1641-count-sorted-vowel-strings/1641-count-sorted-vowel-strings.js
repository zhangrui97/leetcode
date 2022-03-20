/**
 * @param {number} n
 * @return {number}
 */
var countVowelStrings = function(n) {
  const oneLetter = 5
  const twoLetter = (n-1) * 10
  const threeLetter = (n-1)*(n-2)/2 * 10
  const fourLetter = (n-1)*(n-2)*(n-3)/6 * 5
  const fiveLetter = (n-1)*(n-2)*(n-3)*(n-4)/24
  return oneLetter + twoLetter + threeLetter + fourLetter + fiveLetter
};