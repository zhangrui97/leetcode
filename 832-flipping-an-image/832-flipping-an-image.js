/**
 * @param {number[][]} image
 * @return {number[][]}
 */
var flipAndInvertImage = function(image) {
  return image.map(r => r.reverse().map(n => n ^ 1))  
};