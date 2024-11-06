file_path = "input/17.txt"

total = 150
containers = []

with open(file_path, "r") as f:
    for line in f:
        container = int(line.strip("\n"))
        containers.append(container)


valid = []

def find(containers, target, start=0, current=[]):
    if target == 0:
        valid.append(list(current))
        return
    elif target < 0:
        return

    for i in range(start, len(containers)):
        current.append(containers[i])
        find(containers, target - containers[i], i + 1, current)
        current.pop() 

find(containers, total, 0, [])
print(len(valid))

valid = sorted(valid, key=len)
count = 0

print(valid[0])

for i in range(1,len(valid)):
    if len(valid[i]) == len(valid[0]):
        print(valid[i])
        count += 1
    else:
        break


print(count)