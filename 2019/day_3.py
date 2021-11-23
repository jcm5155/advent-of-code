f = open("day_3_input.txt", 'r').read().split('\n')
wire_1 = f[0].split(',')
wire_2 = f[1].split(',')

X_dct = {'U': 0, 'D': 0, 'R': 1, 'L': -1}
Y_dct = {'U': 1, 'D': -1, 'R': 0, 'L': 0}

def get_coords(wire):
    x, y, steps = 0, 0, 0
    coord_dct = {}
    for entry in wire:
        direction = entry[0]
        distance = int(entry[1:])
        for _ in range(distance):
            x += X_dct[direction]
            y += Y_dct[direction]
            steps += 1
            if not coord_dct.get((x, y)):
                coord_dct[(x, y)] = steps
    return coord_dct

def get_intersections(c_1, c_2):
    return [(key, c_1[key]+c_2[key]) for key in c_1.keys() if key in c_2.keys()]


coords_1 = get_coords(wire_1)
coords_2 = get_coords(wire_2)
intersections = get_intersections(coords_1, coords_2)

# Part 1
answer_1 = min([abs(x) + abs(y) for ((x, y), _) in intersections])
print(f"part 1 answer = {answer_1}")

# Part 2
answer_2 = min([steps for (_, steps) in intersections])
print(f"part 2 answer = {answer_2}")
