/**
 * @param {string[][]} items
 * @param {string} ruleKey
 * @param {string} ruleValue
 * @return {number}
 */
var countMatches = function(items, ruleKey, ruleValue) {
  const map = key => {
    switch (key) {
      case 'type': return 0
      case 'color': return 1
      case 'name': return 2
      default: 3
    }
  }
  const key = map(ruleKey)
  return items.filter(a => a[key] === ruleValue).length
};