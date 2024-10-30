file_path = "input/02.txt"
f = open(file_path, "r")

wrap_sqft = 0
ribbon_ft = 0

line = f.readline()
while line:
    nums = line.strip("\n").split("x")
    nums = [int(n) for n in nums]

    a = 2 * nums[0] * nums[1]
    b = 2 * nums[1] * nums[2] 
    c = 2 * nums[2] * nums[0]

    wrap_sqft += (a + b + c) 
    wrap_sqft += min(a//2, b//2, c//2)
    
    ribbon_ft += (sum(sorted(nums)[:2]) * 2)
    
    temp = 1
    for n in nums:
        temp *= n 

    ribbon_ft += temp

    line = f.readline()
