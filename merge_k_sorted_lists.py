class ListNode:
     def __init__(self, val=0, next=None):
         self.val = val
         self.next = next
def merge_k_lists(lists):
    aggregate = []
    for l in lists:
        x = l
        while x != None:
            aggregate.append(x.val)
            x = x.next
    aggregate.sort()
    head = tail = None
    for x in aggregate:
        if head == None:
            head = tail = ListNode(x)
        else:
            tail.next = ListNode(x)
            tail = tail.next
    return head
l0 = ListNode(1, ListNode(4, ListNode(5)))
l1 = ListNode(1, ListNode(3, ListNode(4)))
l2 = ListNode(2, ListNode(6))
lr = merge_k_lists([l0, l1, l2])
while lr != None:
    print(lr.val)
    lr = lr.next
