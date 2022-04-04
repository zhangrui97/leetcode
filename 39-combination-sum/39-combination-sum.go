type context struct {
  c []int
  l int
  cmap map[int]int
  t int
  track []int
  space int
  temp [][2]int
  ans [][]int
} 

func (ctx *context)init() {
  ctx.l = len(ctx.c)
  sort.Ints(ctx.c)
  ctx.cmap = make(map[int]int, ctx.l)
  ctx.ans = make([][]int, 0, 150)
  ctx.track = make([]int, 0, ctx.l)
  sum := 0
  for _, v := range ctx.c {
    cnt := ctx.t / v
    ctx.cmap[v] = cnt
    sum += cnt
  }    
  ctx.temp = make([][2]int, 0, ctx.l)
  ctx.space = sum
}

func (ctx *context)addTrack() {
  ans := make([]int, len(ctx.track))
  copy(ans, ctx.track)
  ctx.ans = append(ctx.ans, ans)
}

func (ctx *context)addResult(result [][2]int) {
  ans := make([]int, 0, ctx.space)
  for _, v := range result {
    val := v[0]
    for i := 0; i < v[1]; i++ {
      ans = append(ans, val)
    }
  }
  ctx.ans = append(ctx.ans, ans)
}

func (ctx *context)check(result [][2]int, start, rest int) {
  l := len(ctx.track)
  if l == start {
    return
  }
  x := ctx.track[start]
  if rest < x {
    return
  }
  if l - start == 1 {
    if rest % x == 0 {
      ctx.addResult(append(result, [2]int{x, rest/x}))
    }
    return
  }
  rest -= x
  if rest <= x {
    return
  }
  result = append(result, [2]int{x, 0})
  for i := 0; i < ctx.cmap[x]; i ++ {
    if rest <= 0 {
      break
    }
    result[len(result)-1][1] = i + 1
    ctx.check(result, start + 1, rest)
    rest -= x
  }  
}

func (ctx *context)backtrack(start, sum int) {
  result := ctx.temp[:0]
  if sum > ctx.t { return }
  if sum == ctx.t { 
    ctx.addTrack()
    return
  }
  ctx.check(result, 0, ctx.t)
  for i := start; i < ctx.l; i ++ {
    num := ctx.c[i]
    ctx.track = append(ctx.track, num)
    ctx.backtrack(i + 1, sum + num)
    ctx.track = ctx.track[:len(ctx.track)-1]
  }
}

func combinationSum(candidates []int, target int) [][]int {
  ctx := &context{c: candidates, t: target}
  ctx.init()
  ctx.backtrack(0, 0)
  return ctx.ans
}