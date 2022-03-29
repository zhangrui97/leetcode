func countMatches(items [][]string, ruleKey string, ruleValue string) int {
  ans := 0
  for _, v := range items {
    switch ruleKey {
    case "type": if v[0] == ruleValue { ans++ }
    case "color": if v[1] == ruleValue { ans++ }
    case "name": if v[2] == ruleValue { ans++ }
    }
  }
  return ans
}