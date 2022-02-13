public class Solution {
  public bool IsPalindrome(int x) {
    string str = x.ToString();
    int i = 0;
    int j = str.Length - 1;
    for (;i < j; i++, j--) {
      if (str[i] != str[j]) return false;
    }
    return true;
  }
}