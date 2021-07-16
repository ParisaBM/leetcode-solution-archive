def is_match(s, p):
    states = {0}
    for i in range(len(p)):
        if p[i] == '*':
            continue
        new_states = set()
        if i+1 < len(p) and p[i+1] == '*':
            for state in states:
                j = state
                while True:
                    new_states.add(j)
                    if j == len(s) or s[j] != p[i] and p[i] != '.':
                        break
                    j += 1
        else:
            for state in states:
                if states != len(s) and (p[i] == s[state] or p[i] == '.'):
                    new_states.add(state + 1)
        states = new_states
    return len(s) in states

print(is_match('aa', 'a'))
print(is_match('aa', 'a*'))
print(is_match('ab', '.*'))
print(is_match('aab', 'c*a*b'))
print(is_match('mississippi', 'mis*is*p*.'))
