/**
 * @param {string} s
 * @return {boolean}
 */
var checkString = function(s) {
  const lastA = s.lastIndexOf('a')
  const firstB = s.indexOf('b')
  return (lastA === -1) || (firstB === -1) || (lastA < firstB)
};