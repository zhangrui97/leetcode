import scala.collection.mutable.Map
object Solution {
  private val cache = Map.empty[Tuple2[Int, Int], Boolean]
  private def dp(s: String, si: Int, p: String, pi: Int): Boolean = {
    val sl = s.length
    val pl = p.length
    if (pi >= pl && si >= sl) return true
    if (pi >= pl && si < sl) return false
    if (si >= sl) {
      if ((pl-pi) % 2 != 0) return false
      else {
        for (i <- pi until pl) {
          if ((i-pi)% 2 > 0 && p.charAt(i) != '*') {
            return false
          }
        }
        return true
      }
    }
    
    val pair = (si, pi)
    if (cache contains pair) cache(pair) 
    else {
      val result: Boolean = 
        if (p.charAt(pi) == '.' || p.charAt(pi) == s.charAt(si)) {
          if (pi < (pl-1) && p.charAt(pi+1) == '*')
            dp(s, si, p, pi+2) || dp(s, si+1, p, pi)
          else dp(s, si+1, p, pi+1)
        }
        else {
          if (pi < (pl-1) && p.charAt(pi+1) == '*')
            dp(s, si, p, pi+2)
          else false
        }
      cache.put(pair, result)
      result
    }
  }
  def isMatch(s: String, p: String): Boolean = {
    cache.clear()
    dp(s, 0, p, 0)
  }
}