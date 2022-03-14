/**
 * @param {string} s
 * @return {number}
 */
var countBinarySubstrings = function(s) {
  const len = s.length
  let ans = 0
  let i = s.indexOf('01')
  while (i != -1) {
    ans ++
    let l = i - 1
    let r = i + 2
    while (l >= 0 && r < len && s[l] === '0' && s[r] === '1') {
      ans ++
      l --
      r ++
    }
    i = s.indexOf('01', i+2)
  }
  i = s.indexOf('10')
  while (i != -1) {
    ans ++
    let l = i - 1
    let r = i + 2
    while (l >= 0 && r < len && s[l] === '1' && s[r] === '0') {
      ans ++
      l --
      r ++
    }
    i = s.indexOf('10', i+2)
  }
  return ans
};