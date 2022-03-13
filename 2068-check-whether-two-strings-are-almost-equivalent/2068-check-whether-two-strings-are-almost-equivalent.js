/**
 * @param {string} word1
 * @param {string} word2
 * @return {boolean}
 */
var checkAlmostEquivalent = function(word1, word2) {
  const getMap = word => {
    const map = new Map()
    for (let ch of word) {
      map.set(ch, map.has(ch) ? map.get(ch) + 1 : 1)
    }
    return map
  }
  const map1 = getMap(word1)
  const map2 = getMap(word2)
  for (let ch of new Set([...map1.keys(), ...map2.keys()])) {
    const count1 = map1.has(ch) ? map1.get(ch) : 0
    const count2 = map2.has(ch) ? map2.get(ch) : 0
    console.log(ch, count1, count2)
    if (count1 - count2 > 3 || count2 - count1 > 3)
      return false
  }
  return true
};