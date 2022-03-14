/**
 * @param {string[]} logs
 * @return {number}
 */
var minOperations = function(logs) {
  let level = 0
  for (const log of logs) {
    switch (log) {
      case './': break
      case '../':
        level = Math.max(level - 1, 0)
        break
      default: level ++
    }
  }
  return level
};