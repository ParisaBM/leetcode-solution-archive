from random import randint
def longest_valid_parentheses(s):
    max_len = 0
    stack = []
    left_starts = [None] * len(s)
    right_starts = [None] * len(s)
    for i in range(len(s)-1):
        if s[i:i+2] == '()':
            stack.append((i, i+1))
    while True:
        if stack == []:
            return max_len
        (l, r) = stack.pop()
        if left_starts[l] == None or left_starts[l] < r:
            left_starts[l] = r
        if right_starts[r] == None or right_starts[r] > l:
            right_starts[r] = l
        max_len = max(max_len, r-l+1)
        if 0 <= l-1 and r+1 < len(s) and s[l-1] == '(' and s[r+1] == ')':
            stack.append((l-1, r+1))
        if r+1 < len(s) and left_starts[r+1]:
            stack.append((l, left_starts[r+1]))
        if 0 <= l-1 and right_starts[l-1]:
            stack.append((right_starts, l-1))

s = ''
for i in range(30000):
    if randint(0, 1):
        s += '('
    else:
        s += ')'
print(longest_valid_parentheses(s))
print(longest_valid_parentheses('(()'))
print(longest_valid_parentheses(')()())'))
print(longest_valid_parentheses(''))
print(longest_valid_parentheses('()'*15000))
