/**
 * @param {string} s
 * @return {string}
 */
var replaceDigits = function(s) {
  const shift = (ch, i) => String.fromCharCode(ch.charCodeAt(0) + i)
  return [...s].map((ch, i) => i%2 ? shift(s[i-1], +ch) : ch).join('')
};