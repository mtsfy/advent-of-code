file_path = "input/03.txt"
f = open(file_path, "r")

direction = {
    ">" : (0, 1),   # East (increase y by 1)
    "<" : (0, -1),  # West (decrease y by 1)
    "^" : (-1, 0),  # North (decrease x by 1)
    "v" : (1, 0)    # South (increase x by 1)
}

santa_visited = set()
santa_visited.add((0,0))
sx, sy = 0,0

robo_visited = set()
robo_visited.add((0,0))
rx, ry = 0,0

is_santa_turn = True

for d in f.read():
    temp_x, temp_y = direction[d]
    if is_santa_turn:
        sx += temp_x
        sy += temp_y
        santa_visited.add((sx,sy))
        is_santa_turn = not is_santa_turn
    else:
        rx += temp_x
        ry += temp_y
        robo_visited.add((rx,ry))
        is_santa_turn = True

res = santa_visited.union(robo_visited)

print(len(res))
