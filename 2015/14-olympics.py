file_path = "input/14.txt"

reindeers = {}
total_time = 2503

def calculate_distance(speed, fly_time, rest_time, total_time):
    current_time = 0
    is_resting = False
    total_distance = 0

    while total_time > 0:
        if is_resting:
            total_time -= rest_time
            is_resting = False 
        else:
            total_distance += speed
            current_time += 1

            if current_time == fly_time:
                is_resting = True 
                current_time = 0

            total_time -= 1
    
    return total_distance

def calculate_distances_per_second(speed, fly_time, rest_time, total_time):
    distances = []
    is_resting = False
    total_distance = 0
    
    while total_time > 0:
        if is_resting:
            for _ in range(rest_time):
                if total_time == 0:
                    break
                distances.append(total_distance)
                total_time -= 1
            is_resting = False
        else:
            for _ in range(fly_time):
                if total_time == 0:
                    break
                total_distance += speed
                distances.append(total_distance)
                total_time -= 1
            is_resting = True
    
    return distances

max_distance = float("-inf")

with open(file_path, "r") as f:
    for line in f:
        line = line.strip("\n").split(" ")
        
        reindeer = line[0]
        speed = int(line[3])
        fly_time = int(line[6])
        rest_time = int(line[13])
        
        max_distance = max(max_distance, calculate_distance(speed, fly_time, rest_time, total_time)) # part 1
        reindeers[reindeer] = calculate_distances_per_second(speed, fly_time, rest_time, total_time) # part 2

scores = {name: 0 for name in reindeers}

for second in range(total_time):
    max_distance = max(distances[second] for distances in reindeers.values())

    for reindeer, distances in reindeers.items():
        if distances[second] == max_distance:
            scores[reindeer] += 1

print(max_distance)
print(scores)