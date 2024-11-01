file_path = "input/06.txt"

def lights1():
    houses = [[0] * 1000 for _ in range(1000)]

    with open(file_path, "r") as f:
        for line in f:
            line = line.strip("\n").split(" ")
            
            if len(line) == 4:
                instruction = line[0]
                x1, y1 = int(line[1].split(",")[0]), int(line[1].split(",")[1])
                x2, y2 = int(line[-1].split(",")[0]), int(line[-1].split(",")[1])
            else:
                instruction = line[1]
                x1, y1 = int(line[2].split(",")[0]), int(line[2].split(",")[-1])
                x2, y2 = int(line[-1].split(",")[0]), int(line[-1].split(",")[1])
            
            
            for x in range(x1,x2+1):
                for y in range(y1, y2+1):
                    if instruction == "on":
                        houses[x][y] = 1
                    elif instruction == "off":
                        houses[x][y] = 0
                    else:
                        if houses[x][y] == 0:
                            houses[x][y] = 1
                        else:
                            houses[x][y] = 0

    total = sum(sum(row) for row in houses)
    return total


def lights2():
    houses = [[0] * 1000 for _ in range(1000)]

    with open(file_path, "r") as f:
        for line in f:
            line = line.strip("\n").split(" ")
            
            if len(line) == 4:
                instruction = line[0]
                x1, y1 = int(line[1].split(",")[0]), int(line[1].split(",")[1])
                x2, y2 = int(line[-1].split(",")[0]), int(line[-1].split(",")[1])
            else:
                instruction = line[1]
                x1, y1 = int(line[2].split(",")[0]), int(line[2].split(",")[-1])
                x2, y2 = int(line[-1].split(",")[0]), int(line[-1].split(",")[1])
            
            
            for x in range(x1,x2+1):
                for y in range(y1, y2+1):
                    if instruction == "on":
                        houses[x][y] += 1
                    elif instruction == "off":
                        houses[x][y] = max(houses[x][y] - 1, 0)
                    else:
                        houses[x][y] += 2

    total = sum(sum(row) for row in houses)
    return total

print(lights2())