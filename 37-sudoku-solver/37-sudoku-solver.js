/**
 * @param {character[][]} board
 * @return {void} Do not return anything, modify board in-place instead.
 */
var solveSudoku = function(board) {
  const all = Array.from({length: 9}, (_, i) => '' + (i + 1))
  const rows = Array(9).fill(all).map( a => new Set(a))
  const cols = Array(9).fill(all).map( a => new Set(a))
  const boxes = Array(9).fill(all).map( a => new Set(a))
  const unsolved = []
  const rc2b = ([i, j]) => (~~(i/3))*3+(~~(j/3))
  const getValues = ([i, j]) =>
    all.filter(e => rows[i].has(e) && cols[j].has(e) && boxes[rc2b([i, j])].has(e))
  const set = ([i, j], v) => {
    board[i][j] = v
    rows[i].delete(v)
    cols[j].delete(v)
    boxes[rc2b([i, j])].delete(v)
  }
  const unset = ([i, j], v) => {
    rows[i].add(v)
    cols[j].add(v)
    boxes[rc2b([i, j])].add(v)
    board[i][j] = '.'
  }
  const buildContext = () => {
    for (let i = 0; i < 9; i++) {
      for (let j = 0; j < 9; j++) {
        const ch = board[i][j]
        if (ch === '.') {
          unsolved.push([i, j])
        } else {
          rows[i].delete(ch)
          cols[j].delete(ch)
          boxes[rc2b([i, j])].delete(ch)
        }
      }
    }
    unsolved.sort((i0, i1) => getValues(i1).length-getValues(i0).length)
  }
  const solve = () => {
    if (unsolved.length === 0) {
      return true
    }
    if (unsolved.some(i => getValues(i).length === 0)) return false
    const t = unsolved.pop()
    const values = getValues(t)
    for (let v of values) {
      set(t, v)      
      if (solve()) return true
      unset(t, v)
    }
    unsolved.push(t)
    return false
  }
  buildContext()
  solve()  
};