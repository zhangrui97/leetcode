/**
 * @param {string} s
 * @param {string} t
 * @return {character}
 */
var findTheDifference = function(s, t) {
  if (s === '') return t
  const code = ch => ch.charCodeAt(0) - 96
  return String.fromCharCode([...s, ...t].map(code).reduce((c1, c2) => c1 ^ c2) + 96)
};