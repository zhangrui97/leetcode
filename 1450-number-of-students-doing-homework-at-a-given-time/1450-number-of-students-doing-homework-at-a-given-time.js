/**
 * @param {number[]} startTime
 * @param {number[]} endTime
 * @param {number} queryTime
 * @return {number}
 */
var busyStudent = function(startTime, endTime, queryTime) {
  let ans = 0
  for (let i = 0; i < startTime.length; i++) {
    ans += (queryTime >= startTime[i] && queryTime <= endTime[i])
  }
  return ans
};