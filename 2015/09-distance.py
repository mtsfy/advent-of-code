from itertools import permutations

file_path = "input/09.txt"

distances = {}
cities = set()

with open(file_path, "r") as f:
    for line in f:
        line = line.strip("\n").split(" ")

        start = line[0]
        end = line[2]
        distance = int(line[-1])
        
        distances[(start, end)] = distance 

        cities.add(start)
        cities.add(end)

def get_distance(start, end):
    if (start, end) in distances:
        return distances[(start, end)]
    elif (end, start) in distances:
        return distances[(end, start)]
    else:
        return float('-inf')  # -inf -> part 2

min_distance = float("-inf") # -inf -> part 2

for perm in permutations(list(cities)):
    total = 0
    for i in range(len(perm) - 1):
        total += get_distance(perm[i], perm[i + 1])
    
    min_distance = max(min_distance, total)  # max -> part 2

print(min_distance)
