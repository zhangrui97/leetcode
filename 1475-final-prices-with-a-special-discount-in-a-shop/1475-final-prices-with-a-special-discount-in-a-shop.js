/**
 * @param {number[]} prices
 * @return {number[]}
 */
var finalPrices = function(prices) {
  return prices.map((n, i) => {
    if (i === prices.length - 1) return n
    const discount = prices.slice(i+1).find(el => el <= n)
    return discount ? n - discount : n
  })
};