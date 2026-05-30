class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:

        angs = defaultdict(list)

        for s in strs:
            sortedS = ''.join(sorted(s))
            angs[sortedS].append(s)
        
        return list(angs.values())
            




        