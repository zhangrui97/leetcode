class Solution {
  public int lengthOfLongestSubstring(String s) {
    HashMap<Character, Integer> dp = new HashMap<Character, Integer>(); 
    int last = 0;    
    int result = 0;
    for (int i = 0; i < s.length(); i++) {
      char ch = s.charAt(i);
      if (dp.containsKey(ch) && dp.get(ch) >= last) {
        last = dp.get(ch) + 1;
      } else {
        if (i - last + 1 > result) {
          result = i - last + 1;
        }
      }
      dp.put(ch, i);
    }
    return result;
  }
}