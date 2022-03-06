/**
 * @param {number[][]} points
 * @return {number}
 */
var minTimeToVisitAllPoints = function([[x0, y0], ...t]) {
  if (t.length <= 0) return 0
  const [[x1, y1], ..._] = t
  const dx = Math.abs(x1 - x0)
  const dy = Math.abs(y1 - y0)
  return Math.max(dx, dy) + minTimeToVisitAllPoints(t)
};