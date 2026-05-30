class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        values = {}

        for i, n in enumerate(nums):
            difference = target - n
            if difference in values:
                return [values[difference], i]
            values[n] = i