class Solution:
  def lengthOfLongestSubstring(self, s: str) -> int:
    if len(s) == 0: return 0
    charDict = {}
    lastRepeat = 0
    result = 0
    for i in range(len(s)):
      if s[i] in charDict and charDict[s[i]] >= lastRepeat:
        lastRepeat = charDict[s[i]] + 1
      else:
        result = max(result, i + 1 - lastRepeat)
      charDict[s[i]] = i
    return result