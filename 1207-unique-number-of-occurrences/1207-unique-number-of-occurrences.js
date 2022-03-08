/**
 * @param {number[]} arr
 * @return {boolean}
 */
var uniqueOccurrences = function(arr) {
  const map = new Map()
  arr.forEach(n => map.set(n, map.has(n) ? map.get(n) + 1 : 1))
  const counts = [...map.values()]
  return counts.every((c, i) => counts.indexOf(c) === i)
};