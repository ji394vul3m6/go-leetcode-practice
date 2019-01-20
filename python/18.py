class Solution:
    def fourSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[List[int]bas]
        """
        nums.sort()
        return self.nSum(4, nums, target)
    
    def nSum(self, n, nums, target):
        # print("{}Choose {} from {} to {}".format(" " * (4-n), n, nums, target))
        if n > len(nums):
            return []
        if nums[0] > 0 and nums[0] > target:
            return []
        if nums[len(nums) - 1] < 0 and nums[len(nums) - 1] < target:
            return []
        if n == 2:
            ret = []
            left = 0
            right = len(nums) - 1
            while left < right:
                s = nums[left] + nums[right]
                if s == target:
                    ret.append([nums[left], nums[right]])
                    left += 1
                    while left < right and nums[left] == nums[left-1]:
                        left += 1
                elif s < target:
                    left += 1
                else:
                    right -= 1
            return ret
        ret = []
        finish = {}
        for idx, num in enumerate(nums):
            if finish[num] != True:
                continue
            if num in finish:
                continue
            finish[num] = True
            # print("{}Pick {}".format(" " * (4-n), num))
            result = self.nSum(n-1, nums[idx+1:], target - num)
            for r in result:
                r.insert(0, num)
            ret = ret + result
        return ret
    
    def run(self):
        testcase = [
            (
                [1, 0, -1, 0, -2, 2], 0, [
                    [-1, 0, 0, 1],
                    [-2, -1, 1, 2],
                    [-2, 0, 0, 2],
                ]
            ),
            (
                [0, 0, 0, 0], 0, [
                    [0, 0, 0, 0],
                ]
            ),
            (
                [-3,-2,-1,0,0,1,2,3], 0, [
                    [-3,-2,2,3],[-3,-1,1,3],[-3,0,0,3],[-3,0,1,2],[-2,-1,0,3],[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]
                ]
            ),
            ( [1,-2,-5,-4,-3,3,3,5], -11, [
                [-5,-4,-3,1]
            ])
        ]
        for test in testcase:
            ret = self.fourSum(test[0], test[1])
            test[2].sort()
            ret.sort()
            if str(ret) != str(test[2]):
                print("Check fail with input {}, {}\nexpect\t{}\nget\t{}".format(test[0], test[1], str(test[2]), str(ret)))
            else:
                print("Testcase with input {}, {} success".format(test[0], test[1]))

if __name__ == "__main__":
    s = Solution()
    s.run()