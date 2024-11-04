from itertools import permutations

file_path = "input/13.txt"

pairs = {}
persons = set()

with open(file_path, "r") as f:
    for line in f:
        line = line.strip("\n").split(" ")
        
        person = line[0]
        neighbour = line[-1][:-1]
        impact = line[2]
        units = int(line[3])

        units = -units if impact == "lose" else units

        pairs[(person, neighbour)] = pairs.get((person, neighbour), 0) + units


        pairs[(person, "Me")] = pairs.get((person, "Me"), 0) + 0
        pairs[("Me", person)] = pairs.get(("Me", person), 0) + 0
        pairs[(neighbour, "Me")] = pairs.get((neighbour, "Me"), 0) + 0
        pairs[("Me", neighbour)] = pairs.get(("Me", neighbour), 0) + 0
        
        persons.add(person)
        

persons.add("Me")
def get_happiness(p1, p2):
        total = 0
        total += pairs[(p1, p2)]
        total += pairs[(p2, p1)]

        return total

max_happiness = float("-inf")
optimal_arrangment = None

for perm in permutations(list(persons)):
    total_happiness = 0
    for i in range(len(perm) - 1):
        total_happiness += get_happiness(perm[i], perm[i + 1])
        
    total_happiness += get_happiness(perm[0], perm[-1])
    
    if total_happiness > max_happiness:
        optimal_arrangment = perm
        max_happiness = total_happiness
    

print(max_happiness)
print(optimal_arrangment)







    