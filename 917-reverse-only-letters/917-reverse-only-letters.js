/**
 * @param {string} s
 * @return {string}
 */
var reverseOnlyLetters = function(s) {
  const isLetter = ch => (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
  const a = s.split('')
  let l = 0
  let r = a.length - 1
  while (l < r) {
    while (l < r && !isLetter(a[l])) l++
    while (l < r && !isLetter(a[r])) r--
    const temp = a[l]
    a[l] = a[r]
    a[r] = temp
    l++
    r--
  }
  return a.join('')
};