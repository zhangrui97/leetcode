/**
 * @param {number[]} widths
 * @param {string} s
 * @return {number[]}
 */
var numberOfLines = function(widths, s) {
  let len = 0
  let nol = 1
  for (let ch of s) {
    const width = widths[ch.charCodeAt(0) - 97]
    if (len + width > 100) {
      nol ++
      len = width
    } else len += width
  }
  return [nol, len]
};