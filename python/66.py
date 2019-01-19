class Solution:
    def run(self):
        testcase = [
        ([1, 2, 3], [1, 2, 4]),
        ([9], [1, 0]),
        ([9, 9, 9], [1, 0, 0, 0]),
        ]
        for test in testcase:
            ret = self.plusOne(test[0])
            if str(ret) != str(test[1]):
                print("Check fail with input {}, expect {}, get {}".format(test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))
    def plusOne(self, digits):
        """
        :type digits: List[int]
        :rtype: List[int]
        """
        ret = [x for x in digits]
        idx = len(ret) - 1
        while idx >= 0:
            if ret[idx] + 1 == 10:
                ret[idx] = 0
                idx -= 1
            else:
                ret[idx] += 1
                break
        if idx == -1:
            return [1] + ret
        else:
            return ret
