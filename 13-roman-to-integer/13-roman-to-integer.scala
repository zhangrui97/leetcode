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
    val a = s.map(charToInt)
    var result = 0
    for (i <- 0 until a.length - 1) {
      if (a(i) < a(i+1)) {
        result -= a(i)          
      } else {
        result += a(i)
      }
    }
    result += a(a.length-1)
    return result
  }
}