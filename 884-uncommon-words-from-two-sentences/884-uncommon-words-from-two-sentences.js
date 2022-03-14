/**
 * @param {string} s1
 * @param {string} s2
 * @return {string[]}
 */
var uncommonFromSentences = function(s1, s2) {
  const a1 = s1.split(' ')
  const a2 = s2.split(' ')
  const sub1 = a1.filter(word => a1.indexOf(word) === a1.lastIndexOf(word) && a2.indexOf(word) === -1)
  const sub2 = a2.filter(word => a2.indexOf(word) === a2.lastIndexOf(word) && a1.indexOf(word) === -1)
  return [...sub1, ...sub2]
};