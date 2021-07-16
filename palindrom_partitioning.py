def min_cut(s: str) -> int:
    min_table = []
    for i in range(len(s)):
        if s[:i+1] == s[:i+1][::-1]:
            min_table.append(0)
        else:
            min_table.append(min_table[-1]+1)
            for j in range(1, i):
                if s[j:i+1] == s[j:i+1][::-1]:
                    min_table[-1] = min(min_table[-1], min_table[j-1]+1)
    return min_table[-1]

print(min_cut("aab"))
print(min_cut("a"))
print(min_cut("ab"))
