/**
 * @param {string} s
 * @param {number[]} indices
 * @return {string}
 */
var restoreString = function(s, indices) {
  const result = Array(s.length)
  indices.forEach((i, j) => result[i] = s[j])
  return result.join('')
};