class ListNode:
    def __init__(self, val=0, next=None):
         self.val = val
         self.next = next
def reverse(head):
    previous = None
    current = head
    next = head.next
    while True:
        current.next = previous
        if next == None:
            return current
        previous = current
        current = next
        next = next.next
def reverse_k_group(head: ListNode, k: int) -> ListNode:
    if k == 1:
        return head
    result = head
    previous = None
    current = head
    next = head.next
    previous_group_tail = None
    while True:
        seek = current
        for i in range(k):
            if seek == None:
                if previous_group_tail:
                    previous_group_tail.next = current
                return result
            seek = seek.next
        else:
            group_tail = current
            for i in range(k):
                current.next = previous
                previous = current
                current = next
                if next:
                    next = next.next
            if result == head:
                result = previous
            if previous_group_tail:
                previous_group_tail.next = previous
            previous_group_tail = group_tail

for i in range(1, 6):
    l = ListNode(3, ListNode(4, ListNode(5, ListNode(6))))
    l = reverse_k_group(l, i)
    current = l
    while current:
        print(current.val)
        current = current.next
    print()
