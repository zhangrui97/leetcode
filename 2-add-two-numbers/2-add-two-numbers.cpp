/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
  ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    ListNode* p1 = l1;
    ListNode* p2 = l2;
    ListNode result;
    ListNode* p = &result;
    int carry = 0;
    while (carry || p1 || p2) {
      if (p1 && p2) {
        int sum = p1->val + p2->val + carry;
        p->next = new ListNode(sum%10);
        p1 = p1->next;
        p2 = p2->next;
        p = p->next;
        carry = sum / 10;
      } else if (carry) {
        int sum = carry + (p1 ? p1->val : 0) + (p2 ? p2->val : 0);
        p->next = new ListNode(sum%10);
        if (p1) p1 = p1->next;
        if (p2) p2 = p2->next;
        p = p->next;
        carry = sum / 10;
      } else if (p1) {
        p->next = p1;
        p1 = nullptr;
      } else {
        p->next = p2;
        p2 = nullptr;
      }
    }
    return result.next;    
  }
};