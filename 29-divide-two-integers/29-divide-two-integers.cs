public class Solution {
  public int Divide(int dividend, int divisor) {
    if (divisor == 1) return dividend;
    if (dividend == Int32.MinValue && divisor == -1) return Int32.MaxValue;
    long a = Math.Abs((long)dividend), b = Math.Abs((long)divisor);
    long res = 0;
    for (int x = 31; x >= 0; x--)
      if ((a >> x) >= b) {
          res += 1L << x;
          a -= b << x;
      }
    return (dividend > 0) == (divisor > 0) ? (int)res : (int)-res;
  }
}