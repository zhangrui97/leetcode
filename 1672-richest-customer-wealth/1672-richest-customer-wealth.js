/**
 * @param {number[][]} accounts
 * @return {number}
 */
var maximumWealth = function(accounts) {
  return Math.max(...(accounts.map(a => a.reduce((n1, n2) => n1 + n2))))
};