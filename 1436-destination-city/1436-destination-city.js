/**
 * @param {string[][]} paths
 * @return {string}
 */
var destCity = function(paths) {
  const set = new Set(paths.map(([fst, snd]) => fst))
  for (let [fst, snd] of paths) {
    if (!set.has(snd)) return snd
  }
  return ''
};