/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
  const cache = {};
  for(var i=0; i < nums.length; i++){
    var value = nums[i];
    var compliment = target - value;
    if(cache[compliment] !==undefined){
      return [i,cache[compliment]];
    }
    else cache[value] = i;
  }
};