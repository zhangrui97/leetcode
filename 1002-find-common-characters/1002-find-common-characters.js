/**
 * @param {string[]} words
 * @return {string[]}
 */
var commonChars = function(words) {
  const ans = words.map(word => {
    const map = new Map()
    for (let ch of word) {
      map.set(ch, map.has(ch) ? map.get(ch) + 1 : 1)
    }
    return map
  }).reduce((map1, map2) => {
    const result = new Map()
    for (const [k, v] of map1) {
      if (map2.has(k)) result.set(k, Math.min(v, map2.get(k)))
    }
    return result
  })
  const final = []
  for (const [k, v] of ans) {
    for (let i = 0; i < v; i++) final.push(k)
  }
  return final
};