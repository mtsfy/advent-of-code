import re

file_path = "input/16.txt"

target = {
    "children": 3,
    "cats": 7,
    "samoyeds": 2,
    "pomeranians": 3,
    "akitas": 0,
    "vizslas": 0,
    "goldfish": 5,
    "trees": 3,
    "cars": 2,
    "perfumes": 1,
}

aunts = {}

with open(file_path, "r") as f:
    for line in f:
        # line = ['Sue', '500', 'pomeranians', '10', 'cats', '3', 'vizslas', '5']
        line = re.split(r'[ ,:]+', line.strip())

        number = line[1]

        aunts[number] = 0

        for i in range(2, len(line[0:]), 2):
            if line[i] in target:
                if line[i] in ["cats", "trees"]: # part 2
                    if int(target[line[i]]) < int(line[i+1]):
                        aunts[number] += 1
                elif line[i] in ["pomeranians", "goldfish"]: # part 2
                    if int(target[line[i]]) > int(line[i+1]):
                        aunts[number] += 1
                else:
                    if int(target[line[i]]) == int(line[i+1]):
                        aunts[number] += 1


highest = sorted(aunts.items(), key=lambda x: x[1], reverse=True)[0]
print(highest)