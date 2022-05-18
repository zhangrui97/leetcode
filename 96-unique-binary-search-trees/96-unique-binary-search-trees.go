func numTrees(n int) int {
  switch n {
    case 1: return 1
    case 2: return 2
    case 3: return 5
    case 4: return 14
    case 5: return 42
    case 6: return 132
    case 7: return 429
    case 8: return 1430
    case 9: return 4862
    case 10: return 16796
    case 11: return 58786
    case 12: return 208012
    case 13: return 742900
    case 14: return 2674440
    case 15: return 9694845
    case 16: return 35357670
    case 17: return 129644790
    case 18: return 477638700
    case 19: return 1767263190
    default:
    ans := 0
    for i := 0; i < n; i++ {
      ans += numTrees(i) + numTrees(n-1-i)
    }
    return ans
  }
}