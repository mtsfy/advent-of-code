file_path = "input/18.txt"

rows = 100
cols = 100

grid = [[0] * cols for _ in range(rows)]
new_grid = [[0] * cols for _ in range(rows)]

row = 0
col = 0

with open(file_path, "r") as f:
    for line in f:
        line = list(line.strip("\n"))
        
        for house in line:
            if house == "#":
                grid[row][col] = 1
            else:
                grid[row][col] = 0
            
            if col == len(line) - 1:
                row += 1
                col = 0
            else:
                col += 1

# part 2 (corners)
grid[0][0] = 1
grid[0][cols - 1] = 1
grid[rows - 1][0] = 1
grid[rows - 1][cols - 1] = 1

def check_neighbour(row, col):
    total = 0

    neighbor_offsets = [
        (-1, 0), (1, 0), (0, -1), (0, 1),  # Up, Down, Left, Right
        (-1, -1), (-1, 1), (1, -1), (1, 1) # Diagonals
    ]
    
    neighbors = []
    for offset in neighbor_offsets:
        new_row = row + offset[0] 
        new_col = col + offset[1] 
        
        if 0 <= new_row < cols and 0 <= new_col < rows:
            if grid[new_row][new_col] == 1:
                total += 1
            neighbors.append((new_row, new_col))

    return total, neighbors


for i in range(100):
    for row in range(len(grid)):
        for col in range(len(grid[row])):
            current = grid[row][col]

            total, neighbors = check_neighbour(row, col)
            if current: # if on (total_on == 2 or total_on == 3 -> on else off )
                if total == 2 or total == 3:
                    new_grid[row][col] = 1
                else:
                    new_grid[row][col] = 0
            else: # if off (total_on == 3 -> on else off)
                if total == 3:
                    new_grid[row][col] = 1
                else:
                    new_grid[row][col] = 0
    
    # part 2 (corners)
    new_grid[0][0] = 1
    new_grid[0][cols - 1] = 1
    new_grid[rows - 1][0] = 1
    new_grid[rows - 1][cols - 1] = 1
    
    grid = [row[:] for row in new_grid] # deepcopy

total = sum(sum(row) for row in grid)
print(total)
