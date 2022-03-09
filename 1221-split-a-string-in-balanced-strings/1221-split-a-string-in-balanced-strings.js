/**
 * @param {string} s
 * @return {number}
 */
var balancedStringSplit = function(s) {
  const changed = [...s].map(ch => ch === 'R' ? 1 : -1)
  let balance = 0
  let count = 0
  changed.forEach(n => {
    balance += n
    count += balance === 0
  })
  return count
};