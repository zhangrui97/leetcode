/**
 * @param {string} s
 * @return {string}
 */
var longestNiceSubstring = function(s) {
  const findNiceSubstring = str => {
    if (str.length < 2) return []
    const subs = str.split('').map(ch => {
      if (ch >= 'a' && ch <= 'z') {
        return str.indexOf(ch.toUpperCase()) === -1 ? ' ' : ch
      } else {
        return str.indexOf(ch.toLowerCase()) === -1 ? ' ' : ch
      }
    }).join('').split(' ')
    if (subs.length > 1) {
      return [].concat(...subs.map(findNiceSubstring))
    } else return [str]
  }
  const subs = findNiceSubstring(s)
  let ans = ''
  for (let sub of subs) {
    if (sub.length > ans.length) {
      ans = sub
    }
  }
  return ans
};