/**
 * @param {string} s
 * @return {string}
 */
var sortSentence = function(s) {
  return s.split(' ').sort((s1, s2) => (+s1.at(-1)) - (+s2.at(-1))).map(s => s.substring(0, s.length-1)).join(' ')
};