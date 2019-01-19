# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None
    def __str__(self):
        cur = self
        ret = []
        while cur != None:
            ret.append(str(cur.val))
            cur = cur.next
        return ','.join(ret)

def CreateNewLinkList(inputs):
    if len(inputs) <= 0:
        return None
    ret = ListNode(inputs[0])
    cur = ret
    for i in range(1, len(inputs)):
        cur.next = ListNode(inputs[i])
        cur = cur.next
    return ret

class Solution:
    def middleNode(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        cur = head
        ret = head
        while cur.next != None:
            ret = ret.next
            cur = cur.next
            if cur.next != None:
                cur = cur.next
        return ret
    
    def run(self):
        testcases = [
            (CreateNewLinkList([1,2,3,4,5]),CreateNewLinkList([3,4,5])),
            (CreateNewLinkList([1,2,3,4,5,6]), CreateNewLinkList([4, 5, 6])),
        ]
        for test in testcases:
            ret = self.middleNode(test[0])
            if str(ret) != str(test[1]):
                print("Check fail with input {}, expect {}, get {}".format(test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))

if __name__ == "__main__":
    s = Solution()
    s.run()