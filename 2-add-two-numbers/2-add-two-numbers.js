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
  const result = new ListNode()
  let p1 = l1
  let p2 = l2
  let carry = 0
  let last = result
  while (p1 || p2 || carry) {
    let sum = 0
    if (p1) {
      sum += p1.val
      p1 = p1.next
    }
    if (p2) {
      sum += p2.val
      p2 = p2.next
    }
    sum += carry
    last.next = new ListNode(sum % 10)
    last = last.next
    carry = ~~(sum / 10)
  }
  return result.next
}