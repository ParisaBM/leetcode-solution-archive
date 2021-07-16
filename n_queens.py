def solve_n_queens_helper(prefix, n):
    if len(prefix) == n:
        return [prefix]
    solutions = []
    for i in range(n):
        if i in prefix:
            continue
        for j in range(len(prefix)):
            if abs(prefix[j]-i) == len(prefix)-j:
                break
        else:
            solutions += solve_n_queens_helper(prefix+[i], n)
    return solutions

def solve_n_queens(n):
    return [['.'*y+'Q'+'.'*(n-y-1) for y in x] for x in solve_n_queens_helper([], n)]

print(solve_n_queens(3))
print(solve_n_queens(4))
print(solve_n_queens(5))
