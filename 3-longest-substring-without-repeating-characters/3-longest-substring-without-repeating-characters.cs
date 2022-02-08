public class Solution {
  public int LengthOfLongestSubstring(string s) {
    var cache = new Dictionary<char, int>();
    int lastRepeat = 0;
    int result = 0;
    for (int i = 0; i < s.Length; i++) {
      char ch = s[i];
      if (cache.ContainsKey(ch) && cache[ch] >= lastRepeat) {
        lastRepeat = cache[ch] + 1;
      } else {
        result = Math.Max(result, i + 1 - lastRepeat);
      }
      cache[ch] = i;
    }
    
    return result;
  }
}