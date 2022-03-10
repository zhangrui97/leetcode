/**
 * @param {string} rings
 * @return {number}
 */
var countPoints = function(rings) {
  const map = new Map([['R', new Set()], ['G', new Set()], ['B', new Set()]])
  for (let i = 0; i < rings.length/2; i++) {
    map.get(rings[i*2]).add(+rings[i*2+1]) 
  }
  let count = 0
  for (let i = 0; i <= 9; i++) {
    count += map.get('R').has(i) && map.get('G').has(i) && map.get('B').has(i)
   }
  return count
};