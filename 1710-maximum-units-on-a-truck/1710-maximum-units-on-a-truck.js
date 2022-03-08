/**
 * @param {number[][]} boxTypes
 * @param {number} truckSize
 * @return {number}
 */
var maximumUnits = function(boxTypes, truckSize) {
  boxTypes.sort(([b1, a], [b2, b]) => b - a)
  let rest = truckSize
  let ans = 0
  for (let [boxes, units] of boxTypes) {
    if (rest >= boxes) {
      ans += boxes * units
      rest -= boxes
    } else {
      ans += rest * units
      rest = 0
      return ans
    }
  }
  return ans
};