from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groupAnagram = defaultdict(lambda: [])

        for word in strs:
            key = [0] * 26
            for ch in word:
                key[ord(ch) - ord('a')] += 1
            groupAnagram[tuple(key)].append(word)
        
        return list(groupAnagram.values())
        