class Solution:

  def findNumbers(self, nums: List[int]) -> int:
    def isEvenDigits(num) -> bool:
      if num > 999:
        return num <= 9999 or num == 100000
      else:
        return num > 9 and num <= 99
    ans = 0
    for num in nums:
      if isEvenDigits(num):
        ans = ans + 1
    return ans