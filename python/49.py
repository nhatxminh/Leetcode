from typing import List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        result: List[List[str]] = []
        m: dict[str, List[str]] = {}

        for _, string in enumerate(strs):
            sort_string = "".join(sorted(string))
            m[sort_string] = m.get(sort_string, [])
            m[sort_string].append(string)

        for string_list in m.values():
            result.append(string_list)
        return result