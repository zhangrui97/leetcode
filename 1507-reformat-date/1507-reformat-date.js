/**
 * @param {string} date
 * @return {string}
 */
var reformatDate = function(date) {
  const mathMap = {
    "Jan": '01', "Feb": '02', "Mar": '03', "Apr": '04', "May": '05', "Jun": '06', 
    "Jul": '07', "Aug": '08', "Sep": '09', "Oct": '10', "Nov": '11', "Dec": '12'
  }
  const [d, m, y] = date.split(' ')  
  return `${y}-${mathMap[m]}-${isNaN(d[1]) ? '0' + d.substring(0, 1) : d.substring(0, 2)}`
};