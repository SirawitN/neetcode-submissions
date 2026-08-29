from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        results = defaultdict(list)

        for s in strs:
            key = [0]*26
            for ch in s:
                key[ord(ch) - ord('a')] += 1

            results[tuple(key)].append(s)

        return list(results.values())