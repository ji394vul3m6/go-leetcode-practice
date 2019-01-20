
class Solution:
    def run(self):
        testcases = [
            (5, [
                [1],
                [1,1],
                [1,2,1],
                [1,3,3,1],
                [1,4,6,4,1]
            ]),
        ]
        for test in testcases:
            ret = self.generate(test[0])
            if str(ret) != str(test[1]):
                print("Check fail with input {}, expect {}, get {}".format(test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))

    def generate(self, numRows):
        """
        :type numRows: int
        :rtype: List[List[int]]
        """
        if numRows == 0:
            return []
        ret = [[1]]
        for i in range(1, numRows):
            temp = [1]
            for j in range(1, i):
                temp.append(ret[i-1][j-1] + ret[i-1][j])
            temp.append(1)
            ret.append(temp)
        return ret

if __name__ == "__main__":
    s = Solution()
    s.run()