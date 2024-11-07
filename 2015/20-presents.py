input = 36000000

# part 1
total = [0] * 1000000
house = 1
while True:
    for elf in range(house, len(total), house):
        total[elf] += house * 10
    
    if total[house] >= input:
        break
    
    house += 1
    
print(house)

# part 2
presents = [0] * 1000000  
for elf in range(1, 1000000):
    for house in range(elf, min(elf * 50 + 1, 1000000), elf):
        presents[house] += elf * 11

for house in range(1, 1000000):
    if presents[house] >= input:
        break

print(house)