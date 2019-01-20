class Solution:
    def majorityElement(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        showTime = {}
        for num in nums:
            try:
                showTime[num] += 1
            except:
                showTime[num] = 1
        
        minTime = int(len(nums)/3)
        return [x for x in showTime.keys() if showTime[x] > minTime]
    def run(self):
        testcase = [
            ([3, 2, 3], [3]),
            ([1,1,1,3,3,2,2,2], [1,2]),
        ]
        for test in testcase:
            ret = self.majorityElement(test[0])
            ret.sort()
            test[1].sort()
            if str(ret) != str(test[1]):
                print("Check fail with input {}, expect {}, get {}".format(test[0], test[1], ret))
            else:
                print("Testcase with input {} success".format(test[0]))