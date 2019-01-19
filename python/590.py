
# Definition for a Node.
class Node(object):
    def __init__(self, val, children):
        self.val = val
        self.children = children

class Solution(object):
    def postorder(self, root):
        """
        :type root: Node
        :rtype: List[int]
        """
        ret = []
        if root == None:
          return []
        if root.children == None:
          return [root.val]

        for child in root.children:
            ret += self.postorder(child)
        ret.append(root.val)
        return ret
    def run(self):
        testcases = [(  
          Node(1, [
            Node(3, [
              Node(5, []),
              Node(6, []),
            ]),
            Node(2, []),
            Node(4, [])
          ]), [5, 6, 3, 2, 4, 1]),
          (Node(1, [
            Node(2, []),
            Node(3, []),
          ]), [2, 3, 1])]
          
        for test in testcases:
            ret = self.postorder(test[0])
            if str(ret) != str(test[1]):
                print("Check fail with input {}, expect {}, get {}".format(test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))
