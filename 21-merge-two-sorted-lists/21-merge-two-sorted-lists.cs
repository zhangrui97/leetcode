/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
  public ListNode MergeTwoLists(ListNode list1, ListNode list2) => (list1, list2) switch {
    (null, null) => null,
    var (l1, l2) when l1 == null => l2,
    var (l1, l2) when l2 == null => l1,
    var (l1, l2) when (l1.val < l2.val) => new ListNode(l1.val, MergeTwoLists(l1.next, l2)),
    var (l1, l2) when (l1.val >= l2.val) => new ListNode(l2.val, MergeTwoLists(l1, l2.next))
  };
}
