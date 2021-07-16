def remove_invalid_parenthese(s):
    excess_left = excess_right = 0
    depth = 0
    for c in s:
        if c == '(':
            if depth < 0:
                depth = 1
            else:
                depth += 1
        elif c == ')':
            depth -= 1
            if depth < 0:
                excess_right += 1
    excess_left = max(depth, 0)
    return dedup([y for x in clean_right(s, excess_right, 0) for y in clean_left(x, excess_left)])

def clean_left(s, left_to_delete):
    return [reverse(x) for x in clean_right(reverse(s), left_to_delete, 0)]

def reverse(s):
    r = ''
    for c in s[::-1]:
        if c == '(':
            r += ')'
        elif c == ')':
            r += '('
        else:
            r += c
    return r

def clean_right(s, right_to_delete, depth):
    if right_to_delete == 0:
        return [s]
    elif s == '':
        return []
    elif s[0] == ')' and depth == 0:
        return clean_right(s[1:], right_to_delete-1, 0)
    elif s[0] == ')':
        return clean_right(s[1:], right_to_delete-1, depth) + [')'+x for x in clean_right(s[1:], right_to_delete, depth-1)]
    else:
        return [s[0]+x for x in clean_right(s[1:], right_to_delete, depth+int(s[0]=='('))]

def dedup(l):
    return list(set(l))

print(remove_invalid_parenthese(')('))
print(remove_invalid_parenthese('()'))
print(remove_invalid_parenthese('(()'))
print(remove_invalid_parenthese('))'))
print(remove_invalid_parenthese('())'))
print(remove_invalid_parenthese('()())()'))
print(remove_invalid_parenthese('(a)())()'))
