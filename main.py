def brackets_check(s):
    meetings = 0
    for c in s:
        if c in '{([':
            meetings += 1
        elif c in ')}]':
            meetings -= 1
            if meetings < 0:
                return False
    return meetings == 0
string1="{()}{}()" 
string2="{()}(])"
print(brackets_check(string1))
