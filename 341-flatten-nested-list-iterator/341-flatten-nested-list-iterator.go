/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
func reverse(a []*NestedInteger) {
  for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
      a[i], a[j] = a[j], a[i]
  }
}

type NestedIterator struct {
  data []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
  reverse(nestedList)
  return &NestedIterator{nestedList}
}

func (this *NestedIterator) Next() int {
  last := len(this.data)-1
  val := this.data[last].GetInteger()
  this.data = this.data[:last]
  return val
}

func (this *NestedIterator) HasNext() bool {
  l := len(this.data)
  if l <= 0 {
    return false
  }
  for ; l > 0 && !this.data[len(this.data)-1].IsInteger(); l = len(this.data) {
    list := this.data[l-1].GetList()
    reverse(list)
    this.data = this.data[:l-1]
    this.data = append(this.data, list...)
  }
  return len(this.data) > 0
}