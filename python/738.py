class Solution:
    def monotoneIncreasingDigits(self, N):
        """
        :type N: int
        :rtype: int
        """
        nums = [int(s) for s in str(N)]
        for i in range(len(nums)-1, -1, -1):
          for j in range(i, -1, -1):
            if nums[j] > nums[i]:
              for k in range(i, len(nums)):
                nums[k] = 9
              nums[i-1] -= 1
              break
        
        return int(''.join([str(s) for s in nums]))

    def run(self):
        testcase = [(10, 9), (1234, 1234), (332, 299), (101, 99), (100, 99)]
        for test in testcase:
            ret = self.monotoneIncreasingDigits(test[0])
            if ret != test[1]:
                print("Check fail with input {}, expect {}, get {}".format(test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))
