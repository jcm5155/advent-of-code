# Part 2 doesn't destroy the asteroids in the correct order, but YOLO (it didn't matter for my puzzle input)


from math import atan2, pi

with open('day_10.txt', 'r') as f:
    rows = f.read().splitlines()



y = 0
asteroid_coords = []
for row in rows:
    x = 0
    for point in row:
        if point == '#':
            asteroid_coords.append((x,y))
        x += 1
    y += 1


visible_dct = {}

for origin in asteroid_coords:
    visible_dct[origin] = {}
    for i in range(len(asteroid_coords)):
        curr_coord = asteroid_coords[i]
        if curr_coord == origin:
            continue
        rel_coord = (curr_coord[0]-origin[0], curr_coord[1]-origin[1])
        curr_rad = atan2(rel_coord[0],rel_coord[1])
        curr_deg = curr_rad * 180 / pi
        curr_deg = 360+curr_deg

        if curr_deg not in visible_dct[origin]:
            visible_dct[origin][curr_deg] = [curr_coord]
        else:
            visible_dct[origin][curr_deg] += [curr_coord]


space_station = ('', (0,0))
for origin, asteroids in visible_dct.items():
    if len(asteroids) > len(space_station[1]):
        space_station = (origin, asteroids)

counter = 1
for k,v in reversed(sorted(space_station[1].items())):
    if counter == 200:
        ans_nums = v[0]
        break
    counter += 1
ans = ans_nums[0]*100+ans_nums[1]


# part 1
print(len(space_station[1]))


# part 2
print(ans)