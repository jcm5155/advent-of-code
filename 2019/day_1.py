with open('aoc_go/inputs/day_1.txt', 'r') as f:
    nums = [int(i) for i in f.readlines()]

# Part 1
p1_answer = 0
for i in nums:
    p1_answer += i // 3 - 2
print(p1_answer)

# I forgot how I did part one, but I'm not gonna fix go back and fix it lmao

# Part 2
counter = 0
for i in nums:
    while i // 3 - 2 > 0:
        temp = i // 3 - 2
        counter += temp
        i = temp
print(counter)