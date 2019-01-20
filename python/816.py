class Solution:
    def ambiguousCoordinates(self, S):
        """
        :type S: str
        :rtype: List[str]
        """
        numStr = S[1:len(S)-1]
        # print(numStr)
        ret = []
        pairs = self.insertComma(numStr)
        for (front, end) in pairs:
            fronts = self.insertPoint(front)
            ends = self.insertPoint(end)
            for a in fronts:
                for b in ends:
                    ret.append("({}, {})".format(a, b))
        return ret

    def insertComma(self, numStr):
        ret = []
        for i in range(0, len(numStr)-1):
            front, end = numStr[0:i+1], numStr[i+1:]
            intF = int(front)
            intE = int(end)
            # continue when like '00'
            if intF == 0 and len(front) > 1:
                continue
            if intE == 0 and len(end) > 1:
                continue
            ret.append((front, end))
        return ret

    def insertPoint(self, numStr):
        ret = []
        for j in range(0, len(numStr)):
            if j == len(numStr)-1:
                a = numStr
            else:
                a = numStr[0:j+1]+"."+numStr[j+1:]
            if not self.validNum(a):
                continue
            ret.append(a)
        return ret

    def validNum(self, numStr):
        if "." not in numStr:
            # something like '001 or 00' is not valid
            return (int(numStr) != 0 and numStr[0] != '0') or (int(numStr) == 0 and len(numStr) == 1)
        a, b = numStr.split(".")

        # something like '001' or '00' is not valid
        if (int(a) != 0 and a[0] == '0') or (int(a) == 0 and len(a) > 1):
            return False

        # something like '1.100' or '1.00' is not valid
        temp = list(b)
        temp.reverse()
        if int("".join(temp)) == 0 or temp[0] == '0':
            return False
        return True

    def run(self):
        testcase = [
            (
                "(123)", ["(1, 23)", "(12, 3)", "(1.2, 3)", "(1, 2.3)"],
            ),
            (
                "(00011)", ["(0.001, 1)", "(0, 0.011)"],
            ),
            (
                "(0123)", ["(0, 123)", "(0, 12.3)", "(0, 1.23)",
                           "(0.1, 23)", "(0.1, 2.3)", "(0.12, 3)"],
            ),
            (
                "(100)", ["(10, 0)"],
            )
        ]

        for test in testcase:
            ret = self.ambiguousCoordinates(test[0])
            ret.sort()
            test[1].sort()
            if str(ret) != str(test[1]):
                print("Check fail with input {}, expect {}, get {}".format(
                    test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))


if __name__ == "__main__":
    s = Solution()
    s.run()
