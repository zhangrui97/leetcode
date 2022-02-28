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
  private PriorityQueue<ListNode, int> q = new();
  public ListNode MergeKLists(ListNode[] lists) {
    if (lists.Length == 0) return null;
    foreach (ListNode list in lists) {
      if (list != null)
        q.Enqueue(list, list.val);
    }
    ListNode dummy = new ListNode();
    ListNode p = dummy;
    while (q.Count > 0) {
      p.next = q.Dequeue();
      p = p.next;
      if (p.next != null) {
        q.Enqueue(p.next, p.next.val);
      }
    }
    return dummy.next;
  }
}