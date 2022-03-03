/**
 * @param {character[][]} board
 * @return {boolean}
 */
var isValidSudoku = function(board) {
  const hasDuplication = a => a.some((item, index) => item != '.' && a.indexOf(item) != index)
  const getSubBox = (r, c) =>
    [board[r][c], board[r][c+1], board[r][c+2], 
     board[r+1][c], board[r+1][c+1], board[r+1][c+2], 
     board[r+2][c], board[r+2][c+1], board[r+2][c+2]]
  for (let i = 0; i < 9; i++) {
    if (hasDuplication(board[i]) ||
        hasDuplication(board.map(a => a[i])) ||
        hasDuplication(getSubBox(~~(i / 3) * 3, (i % 3) * 3)))
      return false
  }
  return true
};