x = 123
y = 456
d = (x & y) & 0xFFFF    
e = (x | y) & 0xFFFF    
f = (x << 2) & 0xFFFF   
g = (y >> 2) & 0xFFFF   
h = (~x) & 0xFFFF        
i = (~y) & 0xFFFF      

file_path = "input/07.txt"

bitwise_operators = {
    "AND": lambda x, y: (x & y) & 0xFFFF,
    "OR": lambda x, y: (x | y) & 0xFFFF,
    "XOR": lambda x, y: (x ^ y) & 0xFFFF,
    "NOT": lambda x: (~x) & 0xFFFF,
    "LSHIFT": lambda x, n: (x << n) & 0xFFFF,
    "RSHIFT": lambda x, n: (x >> n) & 0xFFFF,
}

signals = {}
wires = {}

with open(file_path, "r") as f:
    for line in f:
        line = line.strip().split(" ")
        if len(line) == 3:
            value = line[0]
            end = line[-1]
            if end == "b": # part 2
                wires[end] = "956"
                continue
            else:
                wires[end] = value
        elif len(line) == 4:
            operator = line[0]
            start = line[1]
            end = line[-1]
            wires[end] = (operator, start)
        elif len(line) == 5:
            start = line[0]
            operator = line[1]
            second = line[2]
            end = line[-1]
            wires[end] = (start, operator, second)

def get_signal(wire):
    if wire.isdigit():  
        return int(wire)

    if wire in signals: 
        return signals[wire]
    
    instruction = wires[wire]

    if isinstance(instruction, str):  
        signal = get_signal(instruction)
    elif len(instruction) == 2:  
        operator, start = instruction
        signal = bitwise_operators[operator](get_signal(start))
    elif len(instruction) == 3:  
        start, operator, second = instruction
        start_value = get_signal(start)
        second_value = int(second) if second.isdigit() else get_signal(second)
        signal = bitwise_operators[operator](start_value, second_value)
    

    signals[wire] = signal
    return signal


print(wires["b"])
print(get_signal("a"))
