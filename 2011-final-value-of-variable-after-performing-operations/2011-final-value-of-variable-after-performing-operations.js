/**
 * @param {string[]} operations
 * @return {number}
 */
var finalValueAfterOperations = function(operations) {
  return operations.map(i => i[1] === '+' ? 1 : -1).reduce((a, b) => a + b)
};