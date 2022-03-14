/**
 * @param {string} text
 * @param {string} first
 * @param {string} second
 * @return {string[]}
 */
var findOcurrences = function(text, first, second) {
  const a = text.split(' ')
  const l = a.length
  let i = a.indexOf(first)
  const ans = []
  while (i != -1 && i < l - 2) {
    if (a[i+1] == second) {
      ans.push(a[i+2])
    } 
    i = a.indexOf(first, i+1)
  }
  return ans
};