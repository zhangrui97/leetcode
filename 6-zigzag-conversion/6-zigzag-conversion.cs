public class Solution {
  public string Convert(string s, int numRows) {
    int n = s.Length;
    if (numRows == 1 || numRows >= n) return s;    
    StringBuilder sb = new StringBuilder();
    int step = 2 * numRows - 2;
    for (int i = 0; i < numRows; i++) {
      int i0 = i;
      int i1 = step - i;
      if (i%(numRows-1) == 0) i1 = n;
      while (i0 < n || i1 < n) {
        if (i0 < n) {
          sb.Append(s[i0]);
          i0 += step;
        }
        if (i1 < n) {
          sb.Append(s[i1]);
          i1 += step;
        }
      }
    }
    return sb.ToString();
  }
}