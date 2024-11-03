import re

input = "cqjxxyzz"

'''
Straight Sequence: It needs a sequence of three consecutive letters, like "abc" or "xyz".
No Confusing Letters: It should not contain 'i', 'o', or 'l'.
Two Non-Overlapping Pairs: There should be at least two different, non-overlapping pairs of identical letters, like "aa" and "bb".
'''

consecutives = r"abc|bcd|cde|def|efg|fgh|pqr|qrs|rst|stu|tuv|uvw|vwx|wxy|xyz"
forbidden = r"[^iol]"
double = r"(.)\1.*(.)\2"

def generate(current):
    current = list(current)
    last_idx = len(current) - 1

    if current[last_idx] != 'z':
        current[last_idx] = chr(ord(current[last_idx])+1)
        new = "".join(current)
    else:
        new = generate("".join(current[0:last_idx])) + "a"
    
    return new

def main():
    new = input
    is_not_found = True 
    
    while is_not_found:
        new = generate(new)
        if re.search(forbidden, new):
            if re.search(double, new):
                if re.search(consecutives, new):
                    is_not_found = False 
        print(new)

main()
