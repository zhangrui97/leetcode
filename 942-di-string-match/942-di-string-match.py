class Solution:
  def diStringMatch(self, s: str) -> List[int]:
    min = 0
    max = len(s)
    ans = []
    for ch in s:
      if ch == 'I':
        ans.append(min)
        min += 1
      else:
        ans.append(max)
        max -= 1
    ans.append(min)
    return ans