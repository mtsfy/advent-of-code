import json

file_path = "input/12.json"

def parser(data):
    if isinstance(data, dict):
        if "red" in data.values(): # part 2
            return 0
        return sum(parser(item) for item in data.values())
    elif isinstance(data, list):
        return sum(parser(item) for item in data)
    elif isinstance(data, int): 
            return data 
    else: 
        return 0


with open(file_path, "r") as f:
    data = json.load(f)
    keys = data.keys()
    total = 0

    for k in keys:
        total += parser(data[k])
       
print(total)