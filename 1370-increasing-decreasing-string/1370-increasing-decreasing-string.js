/**
 * @param {string} s
 * @return {string}
 */
var sortString = function(s) {
  const a = 'a'.charCodeAt(0)
  const map = Array(26).fill(0)
  for (let ch of s) map[ch.charCodeAt(0) - a]++
  const process = m => {
    if (m.length === 0) return ''
    const next = m.map(([ch, count]) => [ch, count-1])
    return next.map(([ch, c]) => ch).join('') + process(next.filter(([ch, count]) => count > 0).reverse())
  }
  return process(map.map((count, i) => [count, i]).filter(([count, i]) => count > 0).map(([count, i]) => [String.fromCharCode(i + a), count]))
};