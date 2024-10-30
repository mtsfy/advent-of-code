file_path = "input/01.txt"
f = open(file_path, "r")

total = 0
basement = None

for i, line in enumerate(f.read()):
    if line == "(":
        total += 1
    else:
        total -= 1
        if total < 0 and not basement:
            basement = i + 1