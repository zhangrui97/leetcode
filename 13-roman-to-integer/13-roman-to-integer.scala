object Solution {
	def charToInt(ch: Char): Int =
		ch match {
			case 'I' => 1
			case 'V' => 5
			case 'X' => 10
			case 'L' => 50
			case 'C' => 100
			case 'D' => 500
			case 'M' => 1000
			case _ => throw new IllegalArgumentException(s"invalid character - ${ch}")
		}
    
  def romanToInt(s: String): Int = {
    var result = 0
    var last = charToInt(s.charAt(0))
    for (i <- 1 until s.length) {
      val current = charToInt(s.charAt(i))
      if (last < current) {
        result -= last          
      } else {
        result += last
      }
      last = current
    }
    result += last
    return result
  }
}