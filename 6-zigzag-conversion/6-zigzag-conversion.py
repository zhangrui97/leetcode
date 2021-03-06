class Solution:
  def convert(self, s: str, numRows: int) -> str:
    n = len(s)
    if numRows == 1 or numRows >= n: return s
    result = ''
    step = 2*numRows - 2
    for i in range(numRows):
      index1 = i
      index2 = step - i if i % (numRows-1) else n
      while index1 < n or index2 < n:
        if index1 < n: 
          result += s[index1]
          index1 += step
        if index2 < n: 
          result += s[index2]
          index2 += step
    return result