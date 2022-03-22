func fib(n int) int {
  if n < 2 {
    return n
  }
  var a, b = 0, 1  
  for i:=2; i <= n; i++ {
    a, b = b, a+b
  }
  return b
}