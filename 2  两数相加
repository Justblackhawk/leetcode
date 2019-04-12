/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        //     5/2 表示求商数为2，5%2表示取余数为1
        //     1/2 -->  0  ,     2%5 --->  2,
        ListNode dummy(0);
        ListNode* head;
        head= &dummy;
        int carry =0;//保存余数
        while(l1 != NULL || l2!=NULL || carry !=0){
            int sum = (l1?l1->val:0) + (l2?l2->val:0) + carry;
            head->next = new ListNode(sum % 10);
            carry = sum/10;
            head = head->next;
            l1 = l1?l1->next:l1;
            l2 = l2?l2->next:l2;
        }
        return dummy.next;
    }
};
