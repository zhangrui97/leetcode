/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
  const addTwoNumbersHelper = (tup1, tup2, regroup) => {
    if (!tup1 || !tup2) {
      if (!tup1 && !tup2) {
        return regroup ? new ListNode(1, null) : null
      } else if (!tup1) {
        if (regroup) {
          const result = tup2.val + regroup
          return new ListNode(result%10, addTwoNumbersHelper(null, tup2.next, ~~(result/10)))
        } else {
          return tup2
        }
      } else {
        return addTwoNumbersHelper(tup2, tup1, regroup)
      }
    } else {
      const result = tup1.val + tup2.val + regroup
      return new ListNode(result%10, addTwoNumbersHelper(tup1.next, tup2.next, ~~(result/10)))
    }
  }
  
  return addTwoNumbersHelper(l1, l2, 0)
}