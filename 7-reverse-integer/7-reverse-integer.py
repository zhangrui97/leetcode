class Solution:
  def reverse(self, x: int) -> int:
    if x < 0: return -self.reverse(-x)
    result = int(str(x)[::-1])
    return result if -0x80000000 <= result <= 0x7FFFFFFF else 0