class Solution:
    def maximumProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        neg = [x for x in nums if x <= 0]
        pos = [x for x in nums if x >= 0]
        pos.sort()
        neg.sort()
        if len(pos) <= 2 and len(pos) > 0:
            return pos[len(pos)-1]*neg[0]*neg[1]
        elif len(pos) == 0:
            if 0 in nums:
                return 0
            neg.reverse()
            return neg[0]*neg[1]*neg[2]
        
        pos.reverse()
        if len(neg) < 2:
            return pos[0]*pos[1]*pos[2]
        
        return pos[0] * max(pos[1]*pos[2], neg[0]*neg[1])
    def run(self):
        testcase = [
            ([1,2,3], 6),
            ([-4,-3,1, 2, 3], 36),
            ([-4, 1, 2, 3], 6),
            ([-4, -3, -2, -1, 0], 0),
            ([-1, -2, -3, -4, -5], -6),
            ([1,0,100], 0),
            ([1, 2, 0, -1, -2], 4),
        ]
        for test in testcase:
            ret = self.maximumProduct(test[0])
            if ret != test[1]:
                print("Check fail with input {}, expect {}, get {}".format(test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))