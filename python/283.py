class Solution:
    def run(self):
        testcase = [
          ([0, 1, 0, 3, 12], [1, 3, 12, 0, 0]),
          ([0, 0, 0, 3, 12], [3, 12, 0, 0, 0]),
          ([1, 2, 3, 4, 5], [1, 2, 3, 4, 5]),
        ]
        for test in testcase:
            self.moveZeroes(test[0])
            if str(test[0]) != str(test[1]):
                print("Check fail with input {}, expect {}, get {}".format(test[0], test[1], test[0]))
            else:
                print("Testcase with input {} success".format(test[0]))
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: void Do not return anything, modify nums in-place instead.
        """
        writeIdx = 0
        for cur in range(0, len(nums)):
          if nums[cur] != 0:
            nums[writeIdx] = nums[cur]
            writeIdx += 1

        for cur in range(writeIdx, len(nums)):
          nums[cur] = 0