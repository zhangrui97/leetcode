/**
 * @param {string} digits
 * @return {string[]}
 */
var letterCombinations = function(digits) {
  const map = [[],[],
               ['a','b','c'],
               ['d','e','f'],
               ['g','h','i'],
               ['j','k','l'],
               ['m','n','o'],
               ['p','q','r','s'],
               ['t','u','v'],
               ['w','x','y','z']]
  const combine = aoa => {
    if (aoa.length === 0) return []
    if (aoa.length === 1) return aoa[0]
    const [h, ...t] = aoa
    return [].concat(...(combine(t).map(s => [...(h.map(c => c+s))])))
  }
  return combine(digits.split('').map(ch => map[+ch]))
};