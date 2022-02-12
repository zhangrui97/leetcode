class Solution:
  def isPalindrome(self, x: int) -> bool:
    isPal = lambda s : len(s) < 2 or (s[0] == s[-1] and isPal(s[1:-1]))
    return isPal(str(x))