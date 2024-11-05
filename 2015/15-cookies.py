from itertools import product

file_path = "input/15.txt"

ingredients = {}
with open(file_path, "r") as f:
    for line in f:
        line = line.strip("\n").split(" ")

        ingredient = line[0][:-1]

        capacity = int(line[2][:-1])
        durability = int(line[4][:-1])
        flavour = int(line[6][:-1])
        texture = int(line[8][:-1])
        calorie = int(line[10])

        ingredients[ingredient] = [capacity, durability, flavour, texture, calorie]

def calculate_score(amounts, ingredients):
    totals = [0, 0, 0, 0]

    for i, (_, properties) in enumerate(ingredients.items()):
        for j in range(4):
            totals[j] += properties[j] * amounts[i]
    
    totals = [max(0, x) for x in totals]
    score = totals[0] * totals[1] * totals[2] * totals[3]
  
    return score

def calculate_calories(amounts, ingredients): # part 2: calculate calories
    total = 0
    for i, (_, properties) in enumerate(ingredients.items()):
        total += properties[4] * amounts[i]
    return total

max_score = 0

for amounts in product(range(101), repeat=len(ingredients)):
    if sum(amounts) == 100:
        if calculate_calories(amounts, ingredients) == 500: # part 2: 500 calories
            score = calculate_score(amounts, ingredients)
            max_score = max(max_score, score)

print(max_score)