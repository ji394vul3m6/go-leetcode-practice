class Solution:
    def findErrorNums3(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        appear = {}
        dup = -1
        for i in range(0, len(nums)):
            try:
                appear[nums[i]] += 1
                dup = nums[i]
            except:
                appear[nums[i]] = 1
        l = len(nums)
        diff = int(sum(nums) - l*(l+1)/2)
        return [dup, dup - diff]

    def findErrorNums2(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        nums.sort()
        diff = 1
        missing = -1
        dup = -1
        for i in range(0, len(nums)):
            if nums[i] - i != diff:
                if dup == -1 and missing == -1:
                    if nums[i] - i == 0:
                        dup = i
                    if nums[i] - i == 2:
                        missing = i+1
                    diff = nums[i] - i
                else:
                    if dup != -1:
                        return [dup, i]
                    else:
                        return [i+1, missing]
        if dup != -1:
            return [dup, len(nums)]
        else:
            return [len(nums), missing]

    def findErrorNums(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        appear = {}
        for i in range(0, len(nums)):
            try:
                appear[nums[i]] += 1
            except:
                appear[nums[i]] = 1
        dup = -1
        miss = -1
        for i in range(0, len(nums)):
            try:
                val = appear[i+1]
                if val == 2:
                    dup = i+1
            except:
                miss = i+1
            if dup != -1 and miss != -1:
                break
        return [dup, miss]

    def run(self):
        testcase = [
            ([1, 1], [1, 2]),
            ([1, 2, 2, 4], [2, 3]),
            ([1, 2, 2, 4, 5], [2, 3]),
            ([5, 3, 2, 2, 1], [2, 4]),
        ]
        for test in testcase:
            ret = self.findErrorNums(test[0])
            if str(ret) != str(test[1]):
                print("Check fail with input {}, expect {}, get {}".format(
                    test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))
