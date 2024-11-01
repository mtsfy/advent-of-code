file_path = "input/05.txt"
import re

def nice1():
    forbidden = {"ab", "cd", "pq", "xy"}
    res = 0

    with open(file_path, "r") as f:
        for line in f:
            line = line.strip("\n")
            
            v_count = len(re.findall(r'[aeiou]', line))
            
            d_count = len(re.findall(r'(.)\1', line))
            
            if any(forbidden_str in line for forbidden_str in forbidden):
                continue
            
            if v_count >= 3 and d_count >= 1:
                res += 1

    return res

def nice2():
    res = 0

    with open(file_path, "r") as f:
        for line in f:
            word = line.strip("\n")

            if re.search(r'(..).*\1', word) and re.search(r'(.).\1', word):
                res += 1

    return res
