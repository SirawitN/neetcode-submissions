from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = defaultdict(lambda: [])

        for word in strs:
            charMap = [0] * 26

            for char in word:
                charMap[ord(char) - ord('a')] += 1

            groups[tuple(charMap)].append(word)

        ans = [group for group in groups.values()]
        return ans