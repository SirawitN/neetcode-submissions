class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        i, j = 0, len(numbers)-1

        while i<j:
            tmp_sum = numbers[i] + numbers[j]

            if tmp_sum==target:
                break
            elif tmp_sum<target:
                i += 1
            else:
                j -= 1
        
        return [i+1, j+1]
        