class Solution(object):
    def countBinarySubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        count=0
        for i in range(0, len(s)-1):
            if s[i] != s[i+1]:
                j = 0
                while i - j >= 0 and i + 1 + j < len(s):
                    if s[i-j] == s[i] and s[i+1+j] == s[i+1]:
                        count += 1
                    else:
                        break
                    j += 1
        return count
    def run(self):
        testcase = [("00110011", 6), ("10101", 4), ("1010001", 4)]
        for test in testcase:
            ret = self.countBinarySubstrings(test[0])
            if ret != test[1]:
                print("Check fail with input {}, expect {}, get {}".format(test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))


class Solution2(object):
    def countBinarySubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        count=0
        start=0
        current=1
        while current < len(s):
            # print("{}:{} {}".format(s, start, current))
            if s[current] == s[start]:
                current += 1
            else:
                length = 1
                while current+length < len(s) and length < current - start:
                    # print("{}:{} {}".format(s, current, length))
                    if s[current+length] == s[current]:
                        length += 1
                    else:
                        break
                count += length
                start = current
                current = current + length
                # print("count: {}".format(count))

        return count
    def run(self):
        testcase = [("00110011", 6), ("10101", 4), ("1010001", 4)]
        for test in testcase:
            ret = self.countBinarySubstrings(test[0])
            if ret != test[1]:
                print("Check fail with input {}, expect {}, get {}".format(test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))
