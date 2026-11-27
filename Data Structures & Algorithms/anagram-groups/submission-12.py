from collections import defaultdict

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = defaultdict(lambda: [])

        for word in strs:
            key = [0] * 26
            
            for ch in word:
                key[ord(ch.lower()) - ord('a')] += 1

            groups[tuple(key)].append(word)

        return list(groups.values())