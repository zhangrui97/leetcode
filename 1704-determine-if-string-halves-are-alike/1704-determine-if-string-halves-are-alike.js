/**
 * @param {string} s
 * @return {boolean}
 */
var halvesAreAlike = function(s) {
  const vowels = new Set(['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'])
  let i = 0
  let count = 0
  for (; i < s.length / 2; i++) {
    if (vowels.has(s[i])) count++
  }
  for (; i < s.length; i++) {
    if (vowels.has(s[i])) count--
    if (count < 0) return false    
  }
  return count === 0
};