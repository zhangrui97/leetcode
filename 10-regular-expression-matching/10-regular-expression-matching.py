class Solution:
  def dp(self, s: str, si: int, p: str, pi: int) -> bool:
    sl = len(s)
    pl = len(p)
    if (si >= sl and pi >= pl): return True
    if (si < sl and pi >= pl): return False
    if (si >= sl and pi < pl): 
      return (pl-pi)%2 == 0 and all(ch == '*' for ch in p[pi+1::2])
    key = (si, pi)
    if key in self.cache: return self.cache.get(key)
    result = False
    if p[pi] == '.' or s[si] == p[pi]:
      if (pi+1) < pl and p[pi+1] == '*':
        result = self.dp(s, si, p, pi + 2) or self.dp(s, si+1, p, pi)
      else: 
        result = self.dp(s, si+1, p, pi+1)
    else:
      if (pi+1) < pl and p[pi+1] == '*':
        result = self.dp(s, si, p, pi + 2)
      else:
        result = False
    self.cache[key] = result
    return result

  def isMatch(self, s: str, p: str) -> bool:
    self.cache = {}
    return self.dp(s, 0, p, 0)
