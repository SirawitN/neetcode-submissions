class Solution:
    def twoSum(self, nums: List[int], target: int, start: int, stop: int) -> List[List[int]]:
        two_pairs = []
        i, j = start, stop
        
        while i<j:
            twosum = nums[i] + nums[j]
            if twosum == target:
                two_pairs.append([nums[i], nums[j]])
                i += 1
                j -= 1
                continue

            if twosum < target:
                i += 1
            else:
                j -= 1

        return two_pairs

    def threeSum(self, nums: List[int]) -> List[List[int]]:
        sorted_nums = sorted(nums)
        ans = dict()
        n = len(nums)


        for i in range(0, n-2, 1):
            target = -1 * sorted_nums[i]
            pairs = self.twoSum(sorted_nums, target, i+1, n-1)

            if len(pairs) > 0:
                for p in pairs:
                    ans[(sorted_nums[i], p[0], p[1])] = None
            
        return list(ans.keys())