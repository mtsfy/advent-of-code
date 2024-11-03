file_path = "input/08.txt"

total_memo = 0
total_code = 0

old_code = 0
new_code = 0

with open(file_path, "r") as f:
    for line in f:
        line = line.strip("\n")

        
        total_code += len(line) # part 1
        total_memo += len(line[1:-1].encode().decode('unicode_escape')) # part 1
        
        encoded = f'"{line[1:-1]}"' # part 2
        encoded = encoded.replace('\\', '\\\\').replace('"', '\\"') # part 2

        old_code += len(line) # part 2
        new_code += len(encoded) + 2 # part 2



print(new_code)
print(old_code)
print(new_code - old_code)