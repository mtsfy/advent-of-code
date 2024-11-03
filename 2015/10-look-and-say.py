input_sequence = "1113222113"
iterations = 50

for i in range(iterations):
    result = []
    count = 1
    prev_char = input_sequence[0]

    for char in input_sequence[1:]:
        if char == prev_char:
            count += 1
        else:
            result.append(str(count))
            result.append(prev_char)
            prev_char = char
            count = 1

    result.append(str(count))
    result.append(prev_char)
    
    input_sequence = "".join(result)
    print(i)

# print("".join(result))
print(len(result))

