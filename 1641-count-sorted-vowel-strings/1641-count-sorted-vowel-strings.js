/**
 * @param {number} n
 * @return {number}
 */
var countVowelStrings = function(n) {
  const twoLetter = (n-1) * 10
  const threeLetter = twoLetter*(n-2)/2
  const fourLetter = threeLetter*(n-3)/6
  const fiveLetter = fourLetter*(n-4)/20
  return 5 + twoLetter + threeLetter + fourLetter + fiveLetter
};