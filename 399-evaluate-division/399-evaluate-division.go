func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
  parent := make(map[string]string)
  vmap := make(map[string]float64)
  for i := 0; i < len(values); i++ {
    ai, bi := equations[i][0], equations[i][1]
    vb, okb := vmap[bi]
    va, oka := vmap[ai]
    if !(oka || okb) {
      vmap[bi] = 1.0
      vmap[ai] = values[i]
      parent[bi] = bi
      parent[ai] = bi
    } else if okb && !oka {
      vmap[ai] = values[i]*vb
      parent[ai] = parent[bi]
    } else if !okb && oka {
      vmap[bi] = va / values[i]
      parent[bi] = parent[ai]
    } else {
      ap, bp := parent[ai], parent[bi]
      if ap == bp { continue }
      f := vmap[bi]*values[i]/vmap[ai]
      for s, p := range parent {
        if p == ap {
          parent[s] = bp
          vmap[s] *= f
        }
      }
    }
  }
  ans := make([]float64, len(queries))
  for i, v := range queries {
    ci, di := v[0], v[1]
    pc, okc := parent[ci]
    pd, okd := parent[di]
    if okc && okd {
      if pc == pd {
        ans[i] = vmap[ci] / vmap[di]
        continue
      }
    }
    ans[i] = -1.0
  }
  return ans
}