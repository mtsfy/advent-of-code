import random

file_path = "input/19.txt"

input = "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr"

transformations = {}
replacements = []

with open(file_path, "r") as f:
    for i, line in enumerate(f):
        line = line.strip("\n").split(" => ")

        start, end = line

        transformations[start] = transformations.get(start, [])
        transformations[start].append(end)
        replacements.append((start, end))

unique_molecules = set()

for start, ends in transformations.items():
    start_len = len(start)
    for i in range(len(input) - start_len + 1):
        if input[i:i + start_len] == start:
            for end in ends:
                new_molecule = input[:i] + end + input[i + start_len:]
                unique_molecules.add(new_molecule)

print(len(unique_molecules))

reverse_replacements = {end: start for start, end in replacements}

count = 0
molecule = input

while molecule != 'e':
    random_molecule = random.choice(list(reverse_replacements.keys()))
    
    if random_molecule in molecule:
        molecule = molecule.replace(random_molecule, reverse_replacements[random_molecule], 1)
        count += 1

print(count)